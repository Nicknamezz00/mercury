/*
 * MIT License
 *
 * Copyright (c) 2023 Runze Wu
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 *
 */
import { create } from "zustand";
import { userServiceClient } from "@/rpc";
import { User } from "@/types/proto/api/v2/user_service";

interface UserV1Store {
  userMapByUsername: Record<string, User>;
  getOrFetchUserByUsername: (username: string) => Promise<User>;
  getUserByUsername: (username: string) => User;
  updateUser: (user: Partial<User>, updateMask: string[]) => Promise<User>;
}

const requestCache = new Map<string, Promise<any>>();

const useUserV1Store = create<UserV1Store>()((set, get) => ({
  userMapByUsername: {},
  getOrFetchUserByUsername: async (username: string) => {
    const userMap = get().userMapByUsername;
    if (userMap[username]) {
      return userMap[username] as User;
    }
    if (requestCache.has(username)) {
      return await requestCache.get(username);
    }
    const promisedUser = userServiceClient
      .getUser({
        username: username,
      })
      .then(({ user }) => user);
    requestCache.set(username, promisedUser);
    const user = await promisedUser;
    if (!user) {
      throw new Error("user not found");
    }
    requestCache.delete(username);
    userMap[username] = user;
    set(userMap);
    return user;
  },
  getUserByUsername: (username: string) => {
    const userMap = get().userMapByUsername;
    return userMap[username];
  },
  updateUser: async (user: Partial<User>, updateMask: string[]) => {
    const { user: updateUser } = await userServiceClient.updateUser({
      username: user.username,
      user: user,
      updateMask: updateMask,
    });
    if (!updateUser) {
      throw new Error("user not found");
    }
    const userMap = get().userMapByUsername;
    userMap[updateUser.username] = updateUser;
    set(userMap);
    return updateUser;
  },
}));

export default useUserV1Store;

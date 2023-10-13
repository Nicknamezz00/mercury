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
import axios from "axios";
import { Resource } from "@/types/proto/api/v2/resource_service";

export function getSystemStatus() {
  return axios.get<SystemStatus>("/api/v1/status");
}

export function getSystemSetting() {
  return axios.get<SystemSetting[]>("/api/v1/system/setting");
}

export function upsertSystemSetting(systemSetting: SystemSetting) {
  return axios.post<SystemSetting>("/api/v1/system/setting", systemSetting);
}

export function vacuumDatabase() {
  return axios.post("/api/v1/system/vacuum");
}

export function signIn(username: string, password: string) {
  return axios.post("/api/v1/auth/signin", {
    username,
    password,
  });
}

export function signOut() {
  return axios.post("/api/v1/auth/signout");
}

export function createUser(userCreate: UserCreate) {
  return axios.post<User>("/api/v1/user", userCreate);
}

export function getMyselfUser() {
  return axios.get<User>("/api/v1/user/me");
}

export function getUserList() {
  return axios.get<User[]>("/api/v1/user");
}

export function upsertUserSetting(upsert: UserSettingUpsert) {
  return axios.post<UserSetting>(`/api/v1/user/setting`, upsert);
}

export function patchUser(userPatch: UserPatch) {
  return axios.patch<User>(`/api/v1/user/${userPatch.id}`, userPatch);
}

export function deleteUser(userDelete: UserDelete) {
  return axios.delete(`/api/v1/user/${userDelete.id}`);
}

export function getAllMessages(messageFind?: MessageFind) {
  const queryList = [];
  if (messageFind?.offset) {
    queryList.push(`offset=${messageFind.offset}`);
  }
  if (messageFind?.limit) {
    queryList.push(`limit=${messageFind.limit}`);
  }
  if (messageFind?.creatorUsername) {
    queryList.push(`creatorUsername=${messageFind.creatorUsername}`);
  }
  return axios.get<Message[]>(`/api/v1/message/all?${queryList.join("&")}`);
}

export function getMessageList(messageFind?: MessageFind) {
  const queryList = [];
  if (messageFind?.creatorUsername) {
    queryList.push(`creatorUsername=${messageFind.creatorUsername}`);
  }
  if (messageFind?.rowStatus) {
    queryList.push(`rowStatus=${messageFind.rowStatus}`);
  }
  if (messageFind?.pinned) {
    queryList.push(`pinned=${messageFind.pinned}`);
  }
  if (messageFind?.offset) {
    queryList.push(`offset=${messageFind.offset}`);
  }
  if (messageFind?.limit) {
    queryList.push(`limit=${messageFind.limit}`);
  }
  return axios.get<Message[]>(`/api/v1/message?${queryList.join("&")}`);
}

export function getMessageStats(username: string) {
  return axios.get<number[]>(`/api/v1/message/stats?creatorUsername=${username}`);
}

export function getMessageById(id: MessageId) {
  return axios.get<Message>(`/api/v1/message/${id}`);
}

export function createMessage(messageCreate: MessageCreate) {
  return axios.post<Message>("/api/v1/message", messageCreate);
}

export function patchMessage(messagePatch: MessagePatch) {
  return axios.patch<Message>(`/api/v1/message/${messagePatch.id}`, messagePatch);
}

export function pinMessage(messageId: MessageId) {
  return axios.post(`/api/v1/message/${messageId}/organizer`, {
    pinned: true,
  });
}

export function unpinMessage(messageId: MessageId) {
  return axios.post(`/api/v1/message/${messageId}/organizer`, {
    pinned: false,
  });
}

export function deleteMessage(messageId: MessageId) {
  return axios.delete(`/api/v1/message/${messageId}`);
}

export function createResource(resourceCreate: ResourceCreate) {
  return axios.post<Resource>("/api/v1/resource", resourceCreate);
}

export function createResourceWithBlob(formData: FormData) {
  return axios.post<Resource>("/api/v1/resource/blob", formData);
}

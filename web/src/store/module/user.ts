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
import { camelCase } from "lodash-es";
import * as api from "@/helpers/api";
import storage from "@/helpers/storage";
import { getSystemColorTheme } from "@/helpers/utils";
import store, { useAppSelector } from "@/store";
import { setAppearance, setLocale } from "@/store/reducer/global";
import { patchUser, setHost, setUser } from "@/store/reducer/user";

const defaultSetting: Setting = {
  locale: "en",
  appearance: getSystemColorTheme(),
  messageVisibility: "PRIVATE",
};

const defaultLocalSetting: LocalSetting = {
  enableDoubleClickEditing: false,
  dailyReviewTimeOffset: 0,
};

export const convertUserResponse = (user: User): User => {
  const { systemStatus } = store.getState().global;
  const { locale, appearance } = systemStatus.customizedProfile;
  const systemSetting = { locale, appearance };
  const setting: Setting = {
    ...defaultSetting,
    ...systemSetting,
  };
  const { localSetting: storageLocalSetting } = storage.get(["localSetting"]);
  const localSetting: LocalSetting = {
    ...defaultLocalSetting,
    ...storageLocalSetting,
  };
  if (user.userSettingList) {
    for (const userSetting of user.userSettingList) {
      (setting as any)[camelCase(userSetting.key)] = JSON.parse(userSetting.value);
    }
  }

  return {
    ...user,
    setting,
    localSetting,
    createdTimestamp: user.createdTimestamp * 1000,
    updatedTimestamp: user.updatedTimestamp * 1000,
  };
};

export const initialUserState = async () => {
  const { systemStatus } = store.getState().global;
  if (systemStatus.host) {
    store.dispatch(setHost(convertUserResponse(systemStatus.host)));
  }
  const { data } = await api.getMyselfUser();
  if (data) {
    const user = convertUserResponse(data);
    store.dispatch(setUser(user));
    if (user.setting.locale) {
      store.dispatch(setLocale(user.setting.locale));
    }
    if (user.setting.appearance) {
      store.dispatch(setAppearance(user.setting.appearance));
    }
    return user;
  }
};

const doSignIn = async () => {
  const { data: user } = await api.getMyselfUser();
  if (user) {
    store.dispatch(setUser(convertUserResponse(user)));
  } else {
    doSignOut();
  }
  return user;
};

const doSignOut = async () => {
  await api.signOut();
};

export const useUserStore = () => {
  const state = useAppSelector((state) => state.user);

  return {
    state,
    getState: () => {
      return store.getState().user;
    },
    doSignIn,
    doSignOut,
    upsertUserSetting: async (key: string, value: any) => {
      await api.upsertUserSetting({
        key: key as any,
        value: JSON.stringify(value),
      });
      await doSignIn();
    },
    upsertLocalSetting: async (localSetting: LocalSetting) => {
      storage.set({ localSetting });
      store.dispatch(patchUser({ localSetting }));
    },
    patchUser: async (userPatch: UserPatch): Promise<void> => {
      await api.patchUser(userPatch);
      // reload the page if patched user is current user
      if (userPatch.id === store.getState().user.user?.id && userPatch.username) {
        window.location.reload();
      }
    },
    deleteUser: async (userDelete: UserDelete) => {
      await api.deleteUser(userDelete);
    },
  };
};

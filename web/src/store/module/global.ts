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
import * as api from "@/helpers/api";
import storage from "@/helpers/storage";
import { systemServiceClient } from "@/rpc";
import { setAppearance, setGlobalState, setLocale } from "@/store/reducer/global";
import { findLanguageMatch } from "@/utils/i18n";
import store, { useAppSelector } from "../";

export const initialGlobalState = async () => {
  const { locale: storageLocale, appearance: storageAppearance } = storage.get(["locale", "appearance"]);
  const defaultGlobalState = {
    locale: (storageLocale || "en") as Locale,
    appearance: (storageAppearance || "system") as Appearance,
    systemStatus: {
      allowSignUp: false,
      disablePasswordLogin: false,
      disablePublicMessages: false,
      maxUploadSizeMiB: 0,
      autoBackupInterval: 0,
      additionalStyle: "",
      additionalScript: "",
      messageDisplayWithUpdatedTimestamp: false,
      customizedProfile: {
        name: "mercury",
        logoURL: "/logo.png",
        description: "",
        appearance: "system",
        externalURL: "",
      },
    } as SystemStatus,
  };

  const { data } = await api.getSystemStatus();
  if (data) {
    const customizedProfile = data.customizedProfile;
    defaultGlobalState.systemStatus = {
      ...data,
      customizedProfile: {
        name: customizedProfile.name || "mercury",
        logoURL: customizedProfile.logoURL || "/logo.png",
        description: customizedProfile.description,
        locale: customizedProfile.locale || "en",
        appearance: customizedProfile.appearance || "system",
        externalURL: "",
      },
    };
    defaultGlobalState.locale =
      defaultGlobalState.locale || defaultGlobalState.systemStatus.customizedProfile.locale || findLanguageMatch();
    defaultGlobalState.appearance = defaultGlobalState.appearance || defaultGlobalState.systemStatus.customizedProfile.appearance;
  }
  store.dispatch(setGlobalState(defaultGlobalState));
};

export const useGlobalStore = () => {
  const state = useAppSelector((state) => state.global);

  return {
    state,
    getState: () => {
      return store.getState().global;
    },
    getDisablePublicMessages: () => {
      return store.getState().global.systemStatus.disablePublicMessages;
    },
    isDev: () => {
      return state.systemStatus.profile.mode !== "production";
    },
    fetchSystemStatus: async () => {
      const { data: systemStatus } = await api.getSystemStatus();
      const { systemInfo } = await systemServiceClient.getSystemInfo({});
      systemStatus.dbSize = systemInfo?.dbSize || 0;
      store.dispatch(setGlobalState({ systemStatus: systemStatus }));
      return systemStatus;
    },
    setSystemStatus: (systemStatus: Partial<SystemStatus>) => {
      store.dispatch(
        setGlobalState({
          systemStatus: {
            ...state.systemStatus,
            ...systemStatus,
          },
        })
      );
    },
    setLocale: (locale: Locale) => {
      store.dispatch(setLocale(locale));
    },
    setAppearance: (appearance: Appearance) => {
      store.dispatch(setAppearance(appearance));
    },
  };
};

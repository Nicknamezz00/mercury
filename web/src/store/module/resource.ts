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
import { resourceServiceClient } from "@/rpc";
import store, { useAppSelector } from "@/store";
import { useGlobalStore } from "@/store/module/global";
import { patchResource, setResources } from "@/store/reducer/resource";
import { Resource, UpdateResourceRequest } from "@/types/proto/api/v2/resource_service";

export const useResourceStore = () => {
  const state = useAppSelector((state) => state.resource);
  const globalStore = useGlobalStore();
  const maxUploadSizeMiB = globalStore.state.systemStatus.maxUploadSizeMiB;
  // TODO: i18n translate

  return {
    state,
    getState: () => {
      return store.getState().resource;
    },
    async createResource(resourceCreate: ResourceCreate): Promise<Resource> {
      const { data: resource } = await api.createResource(resourceCreate);
      const resourceList = state.resources;
      store.dispatch(setResources([resource, ...resourceList]));
      return resource;
    },
    async createResourceWithBlob(file: File): Promise<Resource> {
      const { name: filename, size } = file;
      if (size > maxUploadSizeMiB * 1024 * 1024) {
        return Promise.reject("maximum upload size is" + maxUploadSizeMiB);
      }

      const formData = new FormData();
      formData.append("file", file, filename);
      const { data: resource } = await api.createResourceWithBlob(formData);
      const resourceList = state.resources;
      store.dispatch(setResources([resource, ...resourceList]));
      return resource;
    },
    async updateResource(request: UpdateResourceRequest): Promise<Resource> {
      const { resource } = await resourceServiceClient.updateResource(request);
      if (!resource) {
        throw new Error("resource is null");
      }
      store.dispatch(patchResource(resource));
      return resource;
    },
  };
};

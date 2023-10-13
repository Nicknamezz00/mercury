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
import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { uniqBy } from "lodash-es";
import { Resource } from "@/types/proto/api/v2/resource_service";

interface State {
  resources: Resource[];
}

const resourceSlice = createSlice({
  name: "resource",
  initialState: {
    resources: [],
  } as State,
  reducers: {
    setResources: (state, action: PayloadAction<Resource[]>) => {
      return {
        ...state,
        resources: action.payload,
      };
    },
    upsertResources: (state, action: PayloadAction<Resource[]>) => {
      return {
        ...state,
        resources: uniqBy([...action.payload, ...state.resources], "id"),
      };
    },
    patchResource: (state, action: PayloadAction<Partial<Resource>>) => {
      return {
        ...state,
        resources: state.resources.map((resource) => {
          if (resource.id === action.payload.id) {
            return {
              ...resource,
              ...action.payload,
            }
          } else {
            return resource;
          }
        }),
      };
    },
  },
});

export const { setResources, upsertResources, patchResource } = resourceSlice.actions;

export default resourceSlice.reducer;

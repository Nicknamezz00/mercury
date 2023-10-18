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
import useCurrentUser from "@/hooks/useCurrentUser";
import { tagServiceClient } from "@/rpc";
import store, { useAppSelector } from "..";
import { deleteTag as deleteTagAction, setTags, upsertTag as upsertTagAction } from "../reducer/tag";

export const useTagStore = () => {
  const state = useAppSelector((state) => state.tag);
  const currentUser = useCurrentUser();

  const getState = () => {
    return store.getState().tag;
  };

  const fetchTags = async () => {
    const { tags } = await tagServiceClient.listTags({
      creatorId: (await currentUser).id,
    });
    store.dispatch(setTags(tags.map((tag) => tag.name)));
  };

  const upsertTag = async (tagName: string) => {
    await tagServiceClient.upsertTag({
      name: tagName,
    });
    store.dispatch(upsertTagAction(tagName));
  };

  const deleteTag = async (tagName: string) => {
    await tagServiceClient.deleteTag({
      tag: {
        name: tagName,
        creatorId: (await currentUser).id,
      },
    });
    store.dispatch(deleteTagAction(tagName));
  };

  return {
    state,
    getState,
    fetchTags,
    upsertTag,
    deleteTag,
  };
};

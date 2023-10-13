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
import { omit } from "lodash-es";
import { create } from "zustand";
import { combine } from "zustand/middleware";
import * as api from "@/helpers/api";
import { DEFAULT_MESSAGE_LIMIT, ROW_STATUE_ARCHIVED, ROW_STATUS_NORMAL } from "@/helpers/consts";
import store, { useAppSelector } from "@/store";
import { createMessage, deleteMessage, patchMessage, upsertMessages } from "@/store/reducer/message";

export const convertMessageResponse = (message: Message): Message => {
  return {
    ...message,
    createdTimestamp: message.createdTimestamp * 1000,
    updatedTimestamp: message.updatedTimestamp * 1000,
    displayTimestamp: message.displayTimestamp * 1000,
  };
};

export const useMessageStore = () => {
  const state = useAppSelector((state) => state.message);
  const messageCacheStore = useMessageCacheStore();
  const fetchMessageById = async (messageId: MessageId) => {
    const { data } = await api.getMessageById(messageId);
    const message = convertMessageResponse(data);
    store.dispatch(upsertMessages([message]));
    return message;
  };

  return {
    state,
    getState: () => {
      return store.getState().message;
    },
    fetchMessages: async (username = "", limit = DEFAULT_MESSAGE_LIMIT, offset = 0) => {
      const messageFind: MessageFind = {
        rowStatus: ROW_STATUS_NORMAL,
        limit: limit,
        offset: offset,
      };
      if (username) {
        messageFind.creatorUsername = username;
      }

      const { data } = await api.getMessageList(messageFind);
      const fetchedMessages = data.map((message) => convertMessageResponse(message));
      store.dispatch(upsertMessages(fetchedMessages));

      for (const m of fetchedMessages) {
        messageCacheStore.setMessageCache(m);
      }
      return fetchedMessages;
    },
    fetchAllMessages: async (limit = DEFAULT_MESSAGE_LIMIT, offset?: number) => {
      const messageFind: MessageFind = {
        rowStatus: ROW_STATUS_NORMAL,
        limit: limit,
        offset: offset,
      };
      const { data } = await api.getAllMessages(messageFind);
      const fetchedMessages = data.map((message) => convertMessageResponse(message));
      store.dispatch(upsertMessages(fetchedMessages));

      for (const m of fetchedMessages) {
        messageCacheStore.setMessageCache(m);
      }
      return fetchedMessages;
    },
    fetchArchivedMessages: async () => {
      const messageFind: MessageFind = {
        rowStatus: ROW_STATUE_ARCHIVED,
      };
      const { data } = await api.getMessageList(messageFind);
      return data.map((m) => {
        return convertMessageResponse(m);
      });
    },
    fetchMessageById,
    getMessageById: async (messageId: MessageId) => {
      for (const m of state.messages) {
        if (m.id === messageId) {
          return m;
        }
      }
      return await fetchMessageById(messageId);
    },
    getLinkedMessages: async (messageId: MessageId): Promise<Message[]> => {
      const regex = new RegExp(`[@(.+?)](${messageId})`);
      return state.messages.filter((m) => m.content.match(regex));
    },
    createMessage: async (messageCreate: MessageCreate) => {
      const { data } = await api.createMessage(messageCreate);
      const message = convertMessageResponse(data);
      store.dispatch(createMessage(message));
      messageCacheStore.setMessageCache(message);
      return message;
    },
    patchMessage: async (messagePatch: MessagePatch): Promise<Message> => {
      const { data } = await api.patchMessage(messagePatch);
      const message = convertMessageResponse(data);
      store.dispatch(patchMessage(omit(message, "pinned")));
      messageCacheStore.setMessageCache(message);
      return message;
    },
    pinMessage: async (MessageId: MessageId) => {
      await api.pinMessage(MessageId);
      store.dispatch(
        patchMessage({
          id: MessageId,
          pinned: true,
        })
      );
    },
    unpinMessage: async (messageId: MessageId) => {
      await api.unpinMessage(messageId);
      store.dispatch(
        patchMessage({
          id: messageId,
          pinned: false,
        })
      );
    },
    deleteMessageById: async (messageId: MessageId) => {
      await api.deleteMessage(messageId);
      store.dispatch(deleteMessage(messageId));
      messageCacheStore.deleteMessageCache(messageId);
    },
  };
};

const useMessageCacheStore = create(
  combine({ messageById: new Map<MessageId, Message>() }, (set, get) => ({
    getState: () => get(),
    getOrFetchMessageById: async (messageId: MessageId) => {
      const message = get().messageById.get(messageId);
      if (message) {
        return message;
      }

      const { data } = await api.getMessageById(messageId);
      const formattedMessage = convertMessageResponse(data);
      set((state) => {
        state.messageById.set(messageId, formattedMessage);
        return state;
      });
      return formattedMessage;
    },
    getMessageById: (messageId: MessageId) => {
      return get().messageById.get(messageId);
    },
    setMessageCache: (message: Message) => {
      set((state) => {
        state.messageById.set(message.id, message);
        return state;
      });
    },
    deleteMessageCache: (messageId: MessageId) => {
      set((state) => {
        state.messageById.delete(messageId);
        return state;
      });
    },
  }))
);

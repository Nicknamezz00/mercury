// MIT License

// Copyright (c) 2023 Runze Wu

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

type Visibility = "PUBLIC" | "PROTECTED" | "PRIVATE";

type MessageId = number;

interface Message {
  id: MessageId;
  creatorUsername: string;
  createdTimestamp: number;
  updatedTimestamp: number;
  rowStatus: RowStatus;
  displayTimestamp: number;
  content: string;
  visibility: Visibility;
  pinned: boolean;
  creatorName: string;
  resourceList: any[];
  relationList: MessageRelation[];
  parent?: Message;
}

interface MessageCreate {
  content: string;
  resourceIdList: ResourceId[];
  relationList: MessageRelationUpsert[];
  visibility?: Visibility;
}

interface MessagePatch {
  id: MessageId;
  createdTimestamp?: number;
  rowStatus?: RowStatus;
  content?: string;
  visibility: Visibility;
  resourceIdList?: ResourceId[];
  relationList?: MessageRelationUpsert[];
}

interface MessageFind {
  creatorUsername?: string;
  rowStatus?: RowStatus;
  pinned?: boolean;
  visibility?: Visibility;
  offset?: number;
  limit?: number;
}
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
import { createChannel, createClientFactory, FetchTransport } from "nice-grpc-web";
import { ResourceServiceDefinition } from "@/types/proto/api/v2/resource_service";
import { SystemServiceDefinition } from "@/types/proto/api/v2/system_service";
import { UserServiceDefinition } from "@/types/proto/api/v2/user_service";

const channel = createChannel(window.location.origin, FetchTransport({ credentials: "include" }));
const clientFactory = createClientFactory();

export const userServiceClient = clientFactory.create(UserServiceDefinition, channel);
export const systemServiceClient = clientFactory.create(SystemServiceDefinition, channel);
export const resourceServiceClient = clientFactory.create(ResourceServiceDefinition, channel);

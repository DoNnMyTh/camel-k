// camel-k: language=groovy dependency=camel-twitter dependency=mvn:com.google.guava:guava:31.1-jre dependency=github:squakez/samplejp:v1.0

/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

from('timer:groovy?period=1000')
    .routeId('groovy')
    .setHeader("m").constant("string!")
    .setBody()
        .simple('Magic${header.m}')
    .to('log:info?showAll=false')

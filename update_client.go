/* Copyright 2025 İrem Kuyucu <irem@digilol.net>
 * Copyright 2025 Laurynas Četyrkinas <laurynas@digilol.net>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package client3xui

import (
	"context"
	"fmt"
	"net/http"
)

// UpdateClientRaw updates client by inboundId and raw JSON string.
// clientId = UUID клиента (идёт в URL)
// rawJson = строка вида {"clients":[{...}]}
func (c *Client) UpdateClientRaw(ctx context.Context, inboundId uint, clientId string, rawJson string) (*ApiResponse, error) {
    req := &AddClientRequest{
        ID:       inboundId,
        Settings: rawJson,
    }

    resp := &ApiResponse{}
    url := fmt.Sprintf("/panel/api/inbounds/updateClient/%s", clientId)

    err := c.Do(ctx, http.MethodPost, url, req, resp)
    if err != nil {
        return nil, err
    }
    if !resp.Success {
        return resp, fmt.Errorf(resp.Msg)
    }
    return resp, nil
}

/*******************************************************************************
*
* Copyright 2019 Stefan Majewsky <majewsky@gmx.net>
*
* This program is free software: you can redistribute it and/or modify it under
* the terms of the GNU General Public License as published by the Free Software
* Foundation, either version 3 of the License, or (at your option) any later
* version.
*
* This program is distributed in the hope that it will be useful, but WITHOUT ANY
* WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
* A PARTICULAR PURPOSE. See the GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License along with
* this program. If not, see <http://www.gnu.org/licenses/>.
*
*******************************************************************************/

package main

import "github.com/matrix-org/gomatrix"

// ResolveRoomAliasResponse is returned by ResolveRoomAlias.
type ResolveRoomAliasResponse struct {
	RoomID  string   `json:"room_id"`
	Servers []string `json:"servers"`
}

// ResolveRoomAlias asks the server to resolve a room alias like
// #foo:example.org into a room ID like !bar:example.org.
func ResolveRoomAlias(client *gomatrix.Client, roomAlias string) (resp *ResolveRoomAliasResponse, err error) {
	urlPath := client.BuildURL("directory", "room", roomAlias)
	err = client.MakeRequest("GET", urlPath, nil, &resp)
	return
}

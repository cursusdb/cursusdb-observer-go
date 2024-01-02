/*
* CursusDB
* Observer for GO
* ******************************************************************
* Originally authored by Alex Gaetano Padula
* Copyright (C) CursusDB
*
* This program is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* This program is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package observer

// Observer is the main CursusDB Observer Node structure
type Observer struct {
	Host      string      // i.e 0.0.0.0
	Port      int         // i.e 7680
	SharedKey string      // Your shared cluster and node key
	Channel   chan string // Observer channel
}

// Listen starts Observer listener
func (ob *Observer) Listen() error {

	return nil
}

// Close closes Observer listener
func (ob *Observer) Close() {

}

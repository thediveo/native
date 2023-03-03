/*
Package native implements translations between numbers in fixed sizes and byte
sequences in native endianess. Think of it as the missing convenience
combination of [bytes.Buffer] and [enconding/binary] for fixed size numbers.

Probably most useful in dealing with [netlink].

[netlink]: https://en.wikipedia.org/wiki/Netlink
*/
package native

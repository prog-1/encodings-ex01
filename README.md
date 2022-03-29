# Base64

See https://en.wikipedia.org/wiki/Base64

In computer programming, Base64 is a group of binary-to-text encoding schemes that represent binary data (more specifically, a sequence of 8-bit bytes) in sequences of 24 bits that can be represented by four 6-bit Base64 digits.

Common to all binary-to-text encoding schemes, Base64 is designed to carry data stored in binary formats across channels that only reliably support text content. Base64 is particularly prevalent on the World Wide Web where one of its uses is the ability to embed image files or other binary assets inside textual assets such as HTML and CSS files.

Base64 is also widely used for sending e-mail attachments. This is required because SMTP—in its original form—was designed to transport 7-bit ASCII characters only. This encoding causes an overhead of 33–36% (33% by the encoding itself; up to 3% more by the inserted line breaks).

## RFC 4648

See https://datatracker.ietf.org/doc/html/rfc4648#section-4

This is the Base64 alphabet uses A-Z, a-z, 0-9 for the first 62 characters, and + with / as the last two characters. Padding is specified as =.

Please note that there are multiple variant tables as shown at https://en.wikipedia.org/wiki/Base64#Variants_summary_table.

## Home Work

1) Implement `base64.Encode` and its unit-tests.
2) Implement `base64.Decode` and its unit-tests.
3) Ensure `encodings-ex01` program prints `true`.

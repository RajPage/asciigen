package ascii

const ASCII_CHARS = " `.-':_,^=;><+!rcz?sLTv)J7(|Fi{C}fI31tlu[neoZ5Yxjya]2ESwqkP6h9d4VpOGbUAKXHm8RD#$Bg0MNWQ%&@"

var BRIGHTNESS = []uint8{0, 3, 6, 9, 11, 14, 17, 20, 23, 26, 29, 31, 34, 37, 40, 43, 46, 49, 51, 54, 57, 60, 63, 66, 69, 71, 74, 77, 80, 83, 86, 89, 91, 94, 97, 100, 103, 106, 109, 111, 114, 117, 120, 123, 126, 129, 131, 134, 137, 140, 143, 146, 149, 151, 154, 157, 160, 163, 166, 169, 171, 174, 177, 180, 183, 186, 189, 191, 194, 197, 200, 203, 206, 209, 211, 214, 217, 220, 223, 226, 229, 231, 234, 237, 240, 243, 246, 249, 251, 255}

// TODO(#2): Allow user to decrease the number of characters used in the ascii art
// Not every image needs 90 characters to represent it.
// I think ideal number will be 61.
// User can use 1-12 which internally will translate to 1, 2, 3, 4, 5, 6, 10, 12, 15, 20, 30 and 60.
// The remaining character is space.

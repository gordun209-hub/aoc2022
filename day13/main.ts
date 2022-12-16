const ff = Deno.readTextFile("input.txt");
// read as json
const input = JSON.parse(await ff);

// .trim() // Remove starting/ending whitespace
// .split("\n"); // Split on newline

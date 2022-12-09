const fs = require('fs');
const solution1 = fs.readFileSync('../input8.txt', 'utf8').
    split('\n').
    map(l => l.split('').map(v => parseInt(v))).
    reduce((acc, row, i, arr) => acc + row.reduce((acc, cur, j) => acc + ((!row.slice(0, j).
        some(t => t >= cur) || !row.slice(j + 1).
            some(t => t >= cur) || !arr.slice(0, i).
                map(r => r[j]).
                some(t => t >= cur) || !arr.slice(i + 1).
                    map(r => r[j]).
                    some(t => t >= cur)) ? 1 : 0), 0), 0);

console.log(solution1)

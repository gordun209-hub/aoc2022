const data = `$ cd /
$ ls
dir a
14848514 b.txt
8504156 c.dat
dir d
$ cd a
$ ls
dir e
29116 f
2557 g
62596 h.lst
$ cd e
$ ls
584 i
$ cd ..
$ cd ..
$ cd d
$ ls
4060174 j
8033020 d.log
5626152 d.ext
7214296 k`


type Dir = {
    name: string
    size: number
    children: Dir[]
} | any


const Limit = 100000
const formatted = data.split("\n")



const getDir = (name: string, size: number, children: Dir[]): Dir => {
    return {
        name,
        size,
        children
    }
}

const parse = (formatted: string[]): Dir => {
    const root = getDir("/", 0, [])
    let currentDir = root
    let currentLevel = 0
    for (let i = 0; i < formatted.length; i++) {
        const line = formatted[i]
        const level = line.split("$").length - 1
        const [size, name] = line.split(" ").filter(Boolean).reverse()
        if (level > currentLevel) {
            currentDir = currentDir.children[currentDir.children.length - 1]
            currentLevel = level
        } else if (level < currentLevel) {
            currentDir = currentDir.parent
            currentLevel = level
        }
        if (size === "dir") {
            const dir = getDir(name, 0, [])
            dir.parent = currentDir
            currentDir.children.push(dir)
        } else {
            currentDir.children.push({
                name,
                size: parseInt(size)
            })
        }
    }
    return root
}







console.log(parse(formatted))

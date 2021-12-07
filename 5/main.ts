import { readFileSync } from 'fs'

let lines = readFileSync("input").toString().split("\r\n")

const dim = 1000

let grid: number[][] = []
for (let i = 0; i < dim; i++) {
  const y = []
  for (let j = 0; j < dim; j++) {
    y.push(0)
  }
  grid.push(y)
}

for (let line of lines) {
  const coords = line.split(" -> ")
  let [x1, y1] = coords[0].split(",")
  .map(coord => parseInt(coord))
  let [x2, y2] = coords[1].split(",")
  .map(coord => parseInt(coord))

  const x = [x1, x2]
  const y = [y1, y2]

  const insert = (x: number, y: number) => {
    grid[x][y]++
  }

  if (x[0] !== x[1] && y[0] !== y[1]) {
    for (let i = x[0], j = y[0]; x[0] < x[1] ? i <= x[1] : i >= x[1]; x[0] < x[1] ? i++ : i--,y[0] < y[1] ? j++ : j--) {
      insert(i, j)
    }
  } else if (x[0] !== x[1]) {
    x.sort((a, b) => a - b)
    y.sort((a, b) => a - b)
    for (let i = x[0]; i <= x[1]; i++) {
      insert(i, y[0])
    }
  } else if (y[0] !== y[1]) {
    x.sort((a, b) => a - b)
    y.sort((a, b) => a - b)
    for (let i = y[0]; i <= y[1]; i++) {
      insert(x[0], i)
    }
  }
}

const soln = grid.reduce((
  previousValue,
  currentValue,
) => previousValue + currentValue.filter(amount => amount >= 2).length, 0)
console.log("min 2:", soln)
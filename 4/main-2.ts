import * as fs from "fs/promises"

async function main() {
  const contents = await fs.readFile("input", {encoding: "utf-8"})
  const lines = contents.split("\r\n")

  const gameNrs = lines.splice(0, 1)[0].split(",")

  let puzzles: string[][][] = []
  let currentPuzzle: string[][] = []

  for (let line of lines) {
    if (!line) {
      if (currentPuzzle.length != 0) {
        puzzles.push(currentPuzzle)
        currentPuzzle = []
      }
      continue
    }
    currentPuzzle.push(line.replaceAll("  ", " ").trim().split(" "))
  }

  const wonIndices: number[] = []
  let lastWonPuzzle: string[][]
  let lastWonGameNr = ""

  for (let gameNr of gameNrs) {
    puzzles = puzzles.map(
      puzzle => puzzle.map(
        rows => rows.map(
          nr => nr === gameNr ? "X" : nr)))

    let foundIndex = null
    while (foundIndex != -1) {
      foundIndex = puzzles.findIndex((puzzle, index) =>
        !wonIndices.includes(index) &&
        (
          puzzle.some(row => 5 === row.reduce((
            previousValue,
            currentValue,
          ) => currentValue === "X" ? previousValue + 1 : previousValue, 0)) ||
          puzzle.some((_, i) => 5 === puzzle.reduce((
            previousValue,
            currentRow,
          ) => currentRow[i] === "X" ? previousValue + 1 : previousValue, 0))
        ),
      )

      if (foundIndex >= 0) {
        wonIndices.push(foundIndex)
        lastWonPuzzle = puzzles[foundIndex].map(rows => rows.map(nr => nr))
        lastWonGameNr = gameNr
      }
    }
  }

  console.log(`score ${calculateScore(lastWonGameNr, lastWonPuzzle)}`)
}

function calculateScore(gameNr: string, puzzle: string[][]) {
  const unmarkedSum = puzzle.reduce((prev, cur) => prev + cur.reduce((prev, cur) => prev + (isNaN(
    parseInt(cur)) ? 0 : parseInt(cur)), 0), 0)
  return parseInt(gameNr) * unmarkedSum
}

main()
module Days.Day1 where

import qualified Data.Set as Set
import qualified Input.Tools as I

executeDay1 :: String -> (Int, Int)
executeDay1 input = (part1 input', part2 input')
  where
    input' = I.toIntList input

part1 :: [Int] -> Int
part1 = sum

part2 :: [Int] -> Int
part2 input = findDuplicate $ scanl (+) 0 (cycle input)

findDuplicate :: [Int] -> Int
findDuplicate = findDuplicate' Set.empty
  where
    findDuplicate' _ [] = error "No value found"
    findDuplicate' a (x : xs)
      | Set.member x a = x
      | otherwise = findDuplicate' (Set.insert x a) xs
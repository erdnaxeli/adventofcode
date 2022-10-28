import Days.Day1
import qualified Input.Tools as I

main :: IO ()
main = print testPart2


testPart2 :: Bool
testPart2 = part2 [7, 7, -2, -7, -4] == 14

testFindDuplicate :: Bool
testFindDuplicate = findDuplicate [1, 2, 3, 2, 18, -1, 2] == 2

testToIntList = I.toIntList "1\n-1\n0\n+1678" == [1, -1, 0, 1678]
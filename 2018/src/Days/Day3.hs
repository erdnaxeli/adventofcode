module Days.Day3 (executeDay3) where

import Text.Regex.PCRE ((=~))

type Id = Int

type Position = (Int, Int)

type Size = (Int, Int)

data Claim = Claim Id Position Size

executeDay3 :: String -> (Int, Int)
executeDay3 input =
  let claims = map readClaim $ lines input
   in (part1 claims, 1)
  where
    readClaim line =
      case line =~ "#(\\d+) @ (\\d+),(\\d+): (\\d+)x(\\d+)" of
        [[_, id', x, y, x', y']] -> Claim (read id') (read x, read y) (read x', read y')
        _ -> error $ "Invalid claim " ++ line

part1 :: [Claim] -> Int
part1 input =
  let maxX = maximum [a + c | Claim _ (a, _) (c, _) <- input]
      maxY = maximum [b + d | Claim _ (_, b) (_, d) <- input]
   in length . filter (>= 2) $ do
        x <- [0 .. maxX]
        y <- [0 .. maxY]
        [length $ filter (contains x y) input]
  where
    contains x y (Claim _ (a, b) (c, d)) =
      a <= x && x <= a + c && b <= y && y <= b + d

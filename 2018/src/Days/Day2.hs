module Days.Day2 (executeDay2) where

import qualified Data.Map as Map
import qualified Data.Set as Set
import Data.Foldable (asum)

executeDay2 :: String -> (Int, String)
executeDay2 input = (part1 input', part2 input')
  where
    input' = lines input

part1 :: [String] -> Int
part1 input =
  let emptyCounter = Map.empty
      addToCounter counter c = Map.insertWith (+) c 1 counter
      countInWords = foldl addToCounter emptyCounter
      extractCount = Set.fromList . Map.elems
      countAllWords = map (extractCount . countInWords)
      sumOccurences = foldl addToCounter
      countOccurences = foldl sumOccurences emptyCounter . countAllWords
      occurences = countOccurences input
   in Map.findWithDefault 0 (2 :: Int) occurences * Map.findWithDefault 0 3 occurences

part2 :: [String] -> String
part2 input =
  let r = asum $ do
        id1 <- input
        id2 <- tail input
        let same = map fst $ filter (uncurry (==)) $ zip id1 id2
        if length same == length id1 - 1
          then return $ Just same
          else return Nothing
   in case r of
        Just x -> x
        _ -> error "No result found"
module Input.Tools (toIntList) where

toIntList :: String -> [Int]
toIntList s = map (read . filter (/= '+')) $ lines s
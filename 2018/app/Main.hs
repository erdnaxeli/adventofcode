module Main (main) where

import Exercise

main :: IO ()
main = do
    r <- executeStr "2"
    print r

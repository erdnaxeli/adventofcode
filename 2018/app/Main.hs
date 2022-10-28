module Main (main) where

import Exercise
import System.Environment (getArgs)

main :: IO ()
main = do
    args <- getArgs
    r <- executeStr $ head args
    print r

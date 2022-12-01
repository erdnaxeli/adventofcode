module Exercise
  ( Exercise,
    executeStr,
  )
where

import qualified Data.ByteString.Lazy.UTF8 as LBSU
import qualified Data.ByteString.UTF8 as BSU
import Data.Time.Clock (getCurrentTime)
import Days.Day1 (executeDay1)
import Days.Day2 (executeDay2)
import Days.Day3 (executeDay3)
import Network.HTTP.Client
  ( createCookieJar,
    insertCookiesIntoRequest,
    newManager,
    receiveSetCookie,
  )
import Network.HTTP.Conduit (tlsManagerSettings)
import Network.HTTP.Simple
  ( defaultRequest,
    getResponseBody,
    getResponseStatusCode,
    httpLBS,
    setRequestHost,
    setRequestManager,
    setRequestPath,
    setRequestPort,
    setRequestSecure,
  )
import Text.Read (readMaybe)
import Web.Cookie (defaultSetCookie, setCookieName, setCookieValue)
import System.Environment (getEnv)

type Day = Int

newtype Exercise = Exercise Day
type Result = Either String (String, String)

executeStr :: String -> IO Result
executeStr day =
  case readMaybe day of
    Nothing -> return $ Left "Invalid day"
    Just x -> do
      inputEither <- getInput x
      return $ do
        input <- inputEither
        execute (Exercise x) input

execute :: Exercise -> String -> Result
execute (Exercise 1) input =
  let (x, y) = executeDay1 input
  in Right (show x, show y)
execute (Exercise 2) input =
  let (x, y) = executeDay2 input
  in Right (show x, y)
execute (Exercise 3) input =
  let (x, y) = executeDay3 input
  in Right (show x, show y)
execute _ _ = Left "Not implemented"

getInput :: Int -> IO (Either String String)
getInput day = do
  manager <- newManager tlsManagerSettings
  now <- getCurrentTime
  sessionValue <- getEnv "AOC_SESSION"
  let setCookie =
        defaultSetCookie
          { setCookieName = BSU.fromString "session",
            setCookieValue = BSU.fromString sessionValue
          }
      request =
        setRequestManager manager
          $ setRequestHost (BSU.fromString "adventofcode.com")
          $ setRequestPath (BSU.fromString $ "/2018/day/" ++ show day ++ "/input")
          $ setRequestPort
            443
          $ setRequestSecure True defaultRequest
      jar = createCookieJar []
      newJar =
        receiveSetCookie setCookie request now True jar
      (cookieRequest, _) = insertCookiesIntoRequest request newJar now
  response <- httpLBS cookieRequest
  case getResponseStatusCode response of
    200 -> return $ Right (LBSU.toString $ getResponseBody response)
    x -> return $ Left ("Error while getting input: " ++ show x)
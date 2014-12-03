module Paths_hackerrank (
    version,
    getBinDir, getLibDir, getDataDir, getLibexecDir,
    getDataFileName, getSysconfDir
  ) where

import qualified Control.Exception as Exception
import Data.Version (Version(..))
import System.Environment (getEnv)
import Prelude

catchIO :: IO a -> (Exception.IOException -> IO a) -> IO a
catchIO = Exception.catch


version :: Version
version = Version {versionBranch = [0,1,0,0], versionTags = []}
bindir, libdir, datadir, libexecdir, sysconfdir :: FilePath

bindir     = "/Users/adamc/Library/Haskell/bin"
libdir     = "/Users/adamc/Library/Haskell/ghc-7.8.3-x86_64/lib/hackerrank-0.1.0.0"
datadir    = "/Users/adamc/Library/Haskell/share/ghc-7.8.3-x86_64/hackerrank-0.1.0.0"
libexecdir = "/Users/adamc/Library/Haskell/libexec"
sysconfdir = "/Users/adamc/Library/Haskell/etc"

getBinDir, getLibDir, getDataDir, getLibexecDir, getSysconfDir :: IO FilePath
getBinDir = catchIO (getEnv "hackerrank_bindir") (\_ -> return bindir)
getLibDir = catchIO (getEnv "hackerrank_libdir") (\_ -> return libdir)
getDataDir = catchIO (getEnv "hackerrank_datadir") (\_ -> return datadir)
getLibexecDir = catchIO (getEnv "hackerrank_libexecdir") (\_ -> return libexecdir)
getSysconfDir = catchIO (getEnv "hackerrank_sysconfdir") (\_ -> return sysconfdir)

getDataFileName :: FilePath -> IO FilePath
getDataFileName name = do
  dir <- getDataDir
  return (dir ++ "/" ++ name)

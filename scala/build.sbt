name := "hackerrank"

version := "1.0"

scalaVersion := "2.11.7"

logLevel in Global := Level.Error

libraryDependencies ++= Seq(
  "org.scalatest" %% "scalatest" % "2.2.4" % "test"
)

resolvers in ThisBuild ++= Seq(
  "Typesafe Repo" at "http://repo.typesafe.com/typesafe/releases/",
  Resolver.sonatypeRepo("snapshots"),
  Resolver.sonatypeRepo("releases")
)
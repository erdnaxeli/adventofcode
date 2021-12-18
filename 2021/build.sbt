val scala3Version = "3.1.0"

lazy val root = project
  .in(file("."))
  .settings(
    name := "Aoc2021",
    version := "0.1.0-SNAPSHOT",
    scalaVersion := scala3Version,
    libraryDependencies += "org.scala-lang.modules" %% "scala-parser-combinators" % "2.1.0",
    run / fork := true
  )

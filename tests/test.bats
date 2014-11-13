#!/usr/bin/env bats

@test "invoking ee site create example.com --wp" {
  run ee site create example.com --wp
  [ "$status" -eq 0 ]
  [ "$output" = "site example.com created with wp option" ]
}

@test "invoking ee site create --wp example.com" {
  run ee site create --wp example.com
  [ "$status" -eq 0 ]
  [ "$output" = "site example.com created with wp option" ]
}

@test "invoking ee site --wp create example.com" {
  run ee site --wp create example.com
  [ "$status" -eq 1 ]
}

@test "invoking ee stack install" {
  run ee stack install
  [ "$status" -eq 0 ]
  [ "$output" = "installed stack" ]
}

@test "invoking ee stack install --nginx" {
  run ee site create --wp example.com
  [ "$status" -eq 0 ]
  [ "$output" = "installed stack : nginx" ]
}

@test "invoking ee stack install --php" {
  run ee stack install --php
  [ "$status" -eq 0 ]
  [ "$output" = "installed stack : php" ]
}

@test "invoking ee stack install --mysql" {
  run ee stack install --mysql
  [ "$status" -eq 0 ]
  [ "$output" = "installed stack : mysql" ]
}

@test "invoking ee stack --nginx install" {
  run ee stack --nginx install
  [ "$status" -eq 1 ]
}
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

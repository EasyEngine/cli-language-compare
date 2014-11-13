#!/usr/bin/env bats

@test "invoking ee site create example.com --wp" {
  run ee site create example.com --wp
  [ "$status" -eq 0 ]
  [ "$output" = "site 'example.com' created with wp option" ]
}

@test "invoking ee site create --wp 1.com" {
  run ee site create --wp 1.com
  [ "$status" -eq 0 ]
  [ "$output" = "site 'example.com' created with wp option" ]
}

@test "invoking ee site --wp create 1.com" {
  run ee site --wp create 1.com
  [ "$status" -eq 1 ]
}

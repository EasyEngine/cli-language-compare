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
  [ "$status" -eq 1 ]
}

@test "invoking ee stack install --web" {
  run ee stack install --web
  [ "$status" -eq 0 ]
  [ "$output" = "Installed stack : web" ]
}

@test "invoking ee stack install --admin" {
  run ee stack install --admin
  [ "$status" -eq 0 ]
  [ "$output" = "Installed stack : admin" ]
}

@test "invoking ee stack install --mail" {
  run ee stack install --mail
  [ "$status" -eq 0 ]
  [ "$output" = "Installed stack : mail" ]
}

@test "invoking ee stack --web install" {
  run ee stack --web install
  [ "$status" -eq 1 ]
}
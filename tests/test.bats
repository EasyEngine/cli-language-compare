#!/usr/bin/env bats

# Test Argument before option
@test "invoking ee site create example.com --wp" {
  run ee site create example.com --wp
  [ "$status" -eq 0 ]
  [ "$output" = "site example.com created with wp option" ]
}

# Test command without argument, option
@test "invoking ee site create" {
  run ee site create 
  [ "$status" -ne 0 ]
}

# Test command without argument
@test "invoking ee site create --wp" {
  run ee site create --wp
  [ "$status" -ne 0 ]
} 

# Test Feature switch (http://click.pocoo.org/3/options/#feature-switches) option for framework
@test "invoking ee site create example.com --wp --php" {
  run ee site create example.com --wp --php
  [ "$status" -ne 0 ]
  [ "$output" = "site example.com created with wp option" ]
}

# Test Option before Argument 
@test "invoking ee site create --wp example.com" {
  run ee site create --wp example.com
  [ "$status" -eq 0 ]
  [ "$output" = "site example.com created with wp option" ]
}

# Test subcommand option run before subcommand
@test "invoking ee site --wp create example.com" {
  run ee site --wp create example.com
  [ "$status" -ne 0 ]
}

# Test command without option where no argument required
@test "invoking ee stack install" {
  run ee stack install
  [ "$status" -eq 0 ]
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
  [ "$status" -ne 0 ]
}

# Test invalid option
@test "invoking ee stack install --mail --test" {
  run ee stack install --mail --test
  [ "$status" -ne 0 ]
}

# Test global options work on subcommands
@test "invoking ee stack install --help" {
  run ee stack install --help
  [ "$status" -eq 0 ]
}
package com.lautaroblasco.habittracker.users.services;

import org.springframework.stereotype.Service;

import com.lautaroblasco.habittracker.users.repositories.UserRepository;

import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class UserService {
    
    private final UserRepository userRepository;
}
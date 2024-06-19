package com.lautaroblasco.habittracker.activities.repositories;

import org.springframework.data.jpa.repository.JpaRepository;
import org.springframework.stereotype.Repository;

import com.lautaroblasco.habittracker.activities.entities.ActivityEntity;

@Repository
public interface ActivityRepository extends JpaRepository<ActivityEntity, Long> {
    
}

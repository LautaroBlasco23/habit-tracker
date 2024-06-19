package com.lautaroblasco.habittracker.activities.entities;

import com.lautaroblasco.habittracker.routines.entity.RoutineEntity;
import com.lautaroblasco.habittracker.users.entities.UserEntity;
import java.util.List;

import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.ManyToMany;
import jakarta.persistence.ManyToOne;
import jakarta.persistence.Table;
import lombok.Builder;
import lombok.Data;

@Entity
@Table(name = "activities")
@Data
@Builder
public class ActivityEntity {
    
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;

    @ManyToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "user_id", referencedColumnName = "id")
    private UserEntity user;

    private String text;

    @ManyToMany(mappedBy = "activities", fetch = FetchType.LAZY)
    private List<RoutineEntity> routines;
}
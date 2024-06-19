package com.lautaroblasco.habittracker.users.entities;

import com.lautaroblasco.habittracker.activities.entities.ActivityEntity;
import com.lautaroblasco.habittracker.auth.entity.AuthEntity;
import com.lautaroblasco.habittracker.routines.entity.RoutineEntity;

import java.util.List;
import java.util.Set;

import jakarta.persistence.CascadeType;
import jakarta.persistence.Column;
import jakarta.persistence.Entity;
import jakarta.persistence.FetchType;
import jakarta.persistence.GeneratedValue;
import jakarta.persistence.GenerationType;
import jakarta.persistence.Id;
import jakarta.persistence.JoinColumn;
import jakarta.persistence.JoinTable;
import jakarta.persistence.ManyToMany;
import jakarta.persistence.OneToOne;
import jakarta.persistence.Table;
import lombok.Builder;
import lombok.Data;

@Entity
@Table(name = "users")
@Data
@Builder
public class UserEntity {
    @Id
    @GeneratedValue(strategy = GenerationType.IDENTITY)
    private Long id;
    
    @OneToOne(fetch = FetchType.LAZY)
    @JoinColumn(name = "auth_id", referencedColumnName = "id")
    private AuthEntity auth;

    private String name;
    private String username;

    @Column(nullable = false)
    private String imageUrl;

    @ManyToMany(fetch = FetchType.LAZY)
    @JoinTable(
        name = "users_routines",
        joinColumns = @JoinColumn(name = "user_id"),
        inverseJoinColumns = @JoinColumn(name = "routine_id")
    )
    private List<RoutineEntity> routines;

    @ManyToMany(fetch = FetchType.LAZY)
    @JoinTable(
        name = "users_activities",
        joinColumns = @JoinColumn(name = "user_id"),
        inverseJoinColumns = @JoinColumn(name = "activity_id")
    )
    List<ActivityEntity> activities;

    @ManyToMany
    @JoinTable(
        name = "users_follows",
        joinColumns = @JoinColumn(name = "followed"),
        inverseJoinColumns = @JoinColumn(name = "follower")
    )
    private Set<UserEntity> followers;

    @ManyToMany(mappedBy = "followers", fetch = FetchType.LAZY, cascade = {CascadeType.REMOVE})
    private Set<UserEntity> following;
}
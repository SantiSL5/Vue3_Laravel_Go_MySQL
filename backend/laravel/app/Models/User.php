<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Foundation\Auth\User as Authenticatable;
use PHPOpenSourceSaver\JWTAuth\Contracts\JWTSubject;


class User extends Authenticatable implements JWTSubject {

    use HasFactory;

    /**
     * The attributes that are mass assignable.
     *
     * @var array
     */
    protected $fillable = [
        'username',
        'email',
        'password',
        'type',
        'image',
        'enabled',
    ];

    /**
     * The attributes that should be hidden for arrays.
     *
     * @var array
     */
    protected $hidden = [
        'password',
    ];

    // public function create($fields) 
    // {
    //     return parent::create([
    //         'username' => $fields['username'],
    //         'email' => $fields['email'],
    //         'password' => bcrypt($fields['password']),
    //         'type' => isset($fields['type']) ? $fields['type'] : 'client',
    //         'image' =>  isset($fields['image']) ? $fields['image'] : 'https://static.productionready.io/images/smiley-cyrus.jpg',
    //         'enabled' =>  isset($fields['enabled']) ? $fields['enabled'] : true,
    //     ]);
    // } 

    public function getJWTIdentifier()
    {
        return $this->getKey();
    }

    public function getJWTCustomClaims()
    {
        return [];
    }

}
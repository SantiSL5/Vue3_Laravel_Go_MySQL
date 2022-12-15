<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;

class DishType extends Model
{
    use HasFactory;
    protected $fillable = [
        'type',
        'photo'
    ];

    // public function dishTypes()
    // {
    //     return $this->hasMany(Menu::class);
    // }
}

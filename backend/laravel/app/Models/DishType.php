<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use App\Http\Models\Dish;

class DishType extends Model
{
    use HasFactory;
    protected $fillable = [
        'type',
        'photo'
    ];

    public function dishTypes()
    {
        return $this->hasMany(Dish::class);
    }
}

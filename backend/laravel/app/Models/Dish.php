<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use App\Http\Models\DishType;

class Dish extends Model
{
    use HasFactory;
    protected $fillable = [
        'dish',
        'type',
        'price',
        'desc',
        'photo'
    ];

    public function dishes()
    {
        return $this->belongsTo(DishType::class);
    }

}   

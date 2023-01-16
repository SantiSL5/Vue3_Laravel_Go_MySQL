<?php

namespace App\Models;

use Illuminate\Database\Eloquent\Factories\HasFactory;
use Illuminate\Database\Eloquent\Model;
use App\Http\Models\User;
use App\Http\Models\Table;

class Reserve extends Model
{
    use HasFactory;
    protected $fillable = [
        'user',
        'table',
        'date',
        'turn',
        'confirmed'
    ];

    public function reserves()
    {
        return $this->belongsTo(User::class);
    }

    public function tables()
    {
        return $this->belongsTo(Table::class);
    }

}   

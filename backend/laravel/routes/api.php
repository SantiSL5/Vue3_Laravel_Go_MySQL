<?php

// use App\Http\Controllers\BookController;
use App\Http\Controllers\TableController;
use App\Http\Controllers\CategoryController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider within a group which
| is assigned the "api" middleware group. Enjoy building your API!
|
*/

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});

Route::prefix('category')->group(function () {
    Route::get('/', [CategoryController::class, 'index']);
    Route::get('/{id}', [CategoryController::class, 'show']);
    Route::post('/', [CategoryController::class, 'store']);
    Route::put('/{id}', [CategoryController::class, 'update']);
    Route::delete('/{id}', [CategoryController::class, 'destroy']);
});

Route::prefix('table')->group(function () {
    Route::get('/', [TableController::class, 'index']);
    Route::get('/{id}', [TableController::class, 'show']);
    Route::post('/', [TableController::class, 'store']);
    Route::put('/{id}', [TableController::class, 'update']);
    Route::delete('/{id}', [TableController::class, 'destroy']);
});

// Route::get('/mesa', [BookController::class, 'prueba']);
// Route::get('/table', [TableController::class, 'prueba']);
// Route::post('/books', [BookController::class, 'addBook']);
// Route::put('/books/{id}', [BookController::class, 'updateBook']);
// Route::delete('/books/{id}', [BookController::class, 'deleteBook']);

<template>
  <div>
    <h2>Books</h2>
    <ul>
      <li v-for="book in books" :key="book.id">{{ book.title }}</li>
    </ul>
    <input v-model="newTitle" placeholder="New book title" />
    <button @click="addBook">Add</button>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
const books = ref([])
const newTitle = ref('')

const fetchBooks = async () => {
  const res = await fetch('/api/books')
  books.value = await res.json()
}

const addBook = async () => {
  if (!newTitle.value) return
  const res = await fetch('/api/books', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ title: newTitle.value })
  })
  const b = await res.json()
  books.value.push(b)
  newTitle.value = ''
}

onMounted(fetchBooks)
</script>

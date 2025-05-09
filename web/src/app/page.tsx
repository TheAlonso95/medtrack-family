import React from 'react';

export default function Home() {
  return (
    <main className="flex min-h-screen flex-col items-center justify-center p-24">
      <h1 className="text-4xl font-bold mb-4">Welcome to the Monorepo Frontend</h1>
      <p className="text-xl">
        This is a Next.js application that is part of a monorepo with Go backend and Kotlin Multiplatform mobile apps.
      </p>
    </main>
  );
}
import type { NextPage } from 'next'
import Head from 'next/head'
import Image from 'next/image'

const Home: NextPage = () => {
  return (
    <div className="flex min-h-screen flex-col items-center justify-center py-2 ">
      <Head>
        <title>Fanzru - Final Project University</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="flex w-full flex-1 flex-col items-center justify-center px-20 text-center">
        <h1 className="text-6xl font-bold">
          Welcome to{' '}
          <a className="text-blue-600" href="https://fanzru.dev">
            Skripsi Fanzru.
          </a>
        </h1>

        <p className="mt-3 mb-2 text-xl">
        {`This website is under construction `}
        </p>
        <code className="rounded-md bg-gray-100 p-2 font-mono text-xs dark:text-black">
            Stay Tune Guys.
        </code>

        <div className="mt-6 flex max-w-4xl flex-wrap items-center justify-around sm:w-full">
         
        </div>
      </main>

      <footer className="flex h-12 w-full items-center justify-center border-t-2 border-dashed">
        <a
          className="flex items-center justify-center gap-2"
          href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
          target="_blank"
          rel="noopener noreferrer"
        >
          Powered by fanzru - Ananda Affan Fattahila
          
        </a>
      </footer>
    </div>
  )
}

export default Home

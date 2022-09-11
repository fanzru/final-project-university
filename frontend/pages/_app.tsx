import '../styles/globals.css'
import type { AppProps } from 'next/app'
import Script from 'next/script';
import { NextSeo } from 'next-seo';
function MyApp({ Component, pageProps }: AppProps) {
  return <div>
    <NextSeo
      title="Fanzru - Final Project University"
      description="Jangan lupa berusahaan walaupun belum tentu berhasil karena menyerah sama saja dengan kalah."
      openGraph={{
        type: 'website',
        url: 'https://skripsi.fanzru.dev/',
        title: 'Fanzru - Final Project University',
        description:
          'Jangan lupa berusaha walaupun belum tentu berhasil karena menyerah sama saja dengan kalah - Fanzru',
        images: [
          {
            url: '/default.jpeg',
          },
        ],
      }}
    />
    <Component {...pageProps} />
  </div>
}

export default MyApp

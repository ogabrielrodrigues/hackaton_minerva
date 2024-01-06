'use client'

import { api_url } from '@/config/api'
import { Answer as AnswerType } from '@/types/answer'
import { useEffect, useState } from 'react'

interface Props {
  answer_id: string
  token: string
}

export function Answer({ answer_id, token }: Props) {
  const [answer, set_answer] = useState<AnswerType | null>(null)

  useEffect(() => {
    async function fetch_data() {
      const response = await fetch(`${api_url}/answer/${answer_id}`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      })

      return (await response.json()) as AnswerType
    }

    fetch_data().then(set_answer)
  }, [answer_id, token])

  return <p>{answer?.content}</p>
}

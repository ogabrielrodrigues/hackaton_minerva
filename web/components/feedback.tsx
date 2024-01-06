import day_js from 'dayjs'
import pt_br from 'dayjs/locale/pt-br'
import { AnswerButton } from './answer_button'
import { GetToken } from '@/hooks/get_token'
import { api_url } from '@/config/api'
import { Answer } from '@/types/answer'
import { Feedback } from '@/types/feedback'

day_js.locale(pt_br)

type Props = Feedback & {
  administrator: boolean
}

export async function Feedback({
  id,
  content,
  administrator,
  answered,
  answer_id,
  sent_at,
}: Props) {
  async function getAnswer() {
    const { token } = await GetToken()

    const response = await fetch(`${api_url}/answer/${answer_id}`, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    })

    const answer = (await response.json()) as Answer

    return answer
  }

  return (
    <div className="flex flex-col bg-zinc-100 border border-zinc-200 p-4 rounded shadow-md gap-2">
      <p className="font-medium">{content}</p>
      <div className="flex items-center justify-between">
        <strong>{day_js(sent_at).format('DD/MM/YYYY')}</strong>
        {administrator && <AnswerButton feedback_id={id} />}
      </div>
      {answered && (
        <div>
          <h2 className="text-primary font-bold">Resposta</h2>
          <p className="font-medium">{(await getAnswer()).content}</p>
        </div>
      )}
    </div>
  )
}

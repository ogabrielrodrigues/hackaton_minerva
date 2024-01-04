import day_js from 'dayjs'
import pt_br from 'dayjs/locale/pt-br'

day_js.locale(pt_br)

interface Props {
  employee_registry: string
  content: string
  sent_at: Date
}

export function Feedback({ content, sent_at }: Props) {
  return (
    <div className="flex flex-col bg-zinc-100 border border-zinc-200 p-4 rounded shadow-md gap-2">
      <p className="font-medium">{content}</p>
      <div className="flex items-center justify-between">
        <strong>{day_js(sent_at).format('DD/MM/YYYY')}</strong>
        <button className="px-2 py-1 bg-primary text-zinc-50 rounded cursor-pointer">
          Responder
        </button>
      </div>
    </div>
  )
}

from flask import Flask, request,jsonify
import subprocess

app = Flask(__name__)


class Result:
    # 传入code和msg
    # code为0表示成功, -1为失败
    def __init__(self, code, msg):
        self.code = code
        self.msg = msg
        self.data = None

    def setData(self, data):
        self.data = data


transcription_result =''


@app.route('/')
def index():
    return "hello world"


def start(audio_file_name):
    import infer
    global transcription_result
    transcription_result = infer.getChunkStr()
    print(transcription_result)


@app.route('/uploadAudio', methods=['POST'])
def uploadAudio():
    # 检查请求中是否包含文件
    if 'file' not in request.files:
        return Result(-1, 'No file part')

    file = request.files['file']

    # 检查用户是否提交了一个空文件
    if file.filename == '':
        return Result(-1, 'No selected file')

    # 检查文件扩展名是否为 .wav
    if not file.filename.endswith('.wav'):
        return Result(-1, 'Invalid file extension')

    # 保存文件到指定路径
    save_path = 'uploads/001.wav'
    file.save(save_path)
    print(f"获取上传的音频: {file.filename}")
    start(save_path)
    return jsonify(code=0, msg='success')


@app.route('/getTranscription')
def getTranscription():
    global transcription_result
    return jsonify(code=0,msg=transcription_result)


if __name__ == '__main__':
    app.run(host='0.0.0.0', port=5000, debug=True)

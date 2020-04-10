<template>
  <div>
    <div class="button-group">
      <p style="text-align:center">{{hint}}</p>
      <van-row type="flex" justify="space-around">
        <van-col span="4">
          <div :class="{'flex-demo':true,'start':isStart,'big-circle':isPc,'small-circle':!isPc}" @click="tryShowAdd">
            Add</div>
        </van-col>
        <van-col span="4">
          <div :class="{'flex-demo':true,'start':isStart,'big-circle':isPc,'small-circle':!isPc}"
            @click="switchConnect">Connect</div>
        </van-col>
        <van-col span="4">
          <div :class="{'flex-demo':true,'start':isStart,'big-circle':isPc,'small-circle':!isPc}" @click="switchDelete">
            Delete</div>
        </van-col>
        <van-col span="4">
          <div :class="{'flex-demo':true,'start':isStart,'big-circle':isPc,'small-circle':!isPc}" @click="deleteAll">
            Delete All</div>
        </van-col>
        <van-col span="4">
          <div :class="{'flex-demo':true,'start':isStart,'big-circle':isPc,'small-circle':!isPc}" @click="recolor">
            Recolor</div>
        </van-col>
        <van-col span="4">
          <div :class="{'flex-demo':true,'start':isStart,'big-circle':isPc,'small-circle':!isPc}" @click="switchReady">
            Start</div>
        </van-col>
      </van-row>
    </div>
    <div>
      <div id="g6"></div>
    </div>

    <van-popup round v-model="showAdd" position="bottom" :style="{ height: '180px' }">
      <div style="width: 95%;background-color:#fff;height:170px;margin:0 auto;border-radius:5px;padding-top:10px;">
        <van-cell-group title="添加节点">
          <van-field v-model="numToAdd" placeholder="请输入 数字 或 [数组] : 比如 1 或 [1,2,3]" />
        </van-cell-group>
        <div style="padding:20px 15px 0px 15px;">
          <van-button type="primary" block @click="addNode">Add</van-button>
        </div>
      </div>
    </van-popup>


  </div>
</template>
<script>
  import G6 from '@antv/g6';
  export default {
    data() {
      return {
        step: 0,
        hint: "",
        numToAdd: "",
        isPc: false,
        showAdd: false,
        isStart: false,
        isReady: false,
        isConnect: false,
        isDelete: false,
        data: null,
        graph: null,
        grid: null,
        tempNode: null,
        clientX: null,
        clientY: null,
        distance: null
      }
    },
    created() {
      if (window.screen.width < 800) {
        this.isPc = false
      } else {
        this.isPc = true
      }

    },
    mounted() {
      this.initG6()
      if (!this.isPc) {
        this.initMobile()
      } else {
        // this.initPC()
      }
    },
    methods: {
      initPC() {
        document.getElementsByTagName("canvas")[0].addEventListener('mousemove', (e) => {
          // console.log(e.type)
          // console.log(e)
        });
        document.getElementsByTagName("canvas")[0].addEventListener('mouseup', (e) => {
          // console.log(e.type)
          // console.log(e)
        });
        document.getElementsByTagName("canvas")[0].addEventListener('mousedown', (e) => {
          // console.log(e.type)
        });
      },
      initMobile() {
        document.getElementsByTagName("canvas")[0].addEventListener('touchstart', (e) => {
          // console.log("触发事件：touchstart");
          // console.log(e.type)
          if (this.clientX == null) {
            this.clientX = e.targetTouches[0].clientX
          }
          if (this.clientY == null) {
            this.clientY = e.targetTouches[0].clientY
          }
          var evt = new MouseEvent("mousedown", {
            bubbles: true,
            cancelable: true,
            view: window,
            screenX: e.targetTouches[0].clientX,
            screenY: e.targetTouches[0].clientY,
            clientX: e.targetTouches[0].clientX,
            clientY: e.targetTouches[0].clientY,
            ctrlKey: false,
            altKey: false,
            shiftKey: false,
            metaKey: false,
            button: 0,
            buttons: 1,
          });
          document.getElementsByTagName("canvas")[0].dispatchEvent(evt)
          e.preventDefault()
        })
        document.getElementsByTagName("canvas")[0].addEventListener('touchend', (e) => {
          // console.log(e.type)
          var evt = new MouseEvent("mouseup", {
            bubbles: true,
            cancelable: true,
            view: window,
            screenX: this.clientX,
            screenY: this.clientY,
            clientX: this.clientX,
            clientY: this.clientY,
            ctrlKey: false,
            altKey: false,
            shiftKey: false,
            metaKey: false,
            button: 0,
            buttons: 1,
          });
          document.getElementsByTagName("canvas")[0].dispatchEvent(evt)
          if (this.clientX !== null) {
            this.clientX = null
          }
          if (this.clientY !== null) {
            this.clientY = null
          }
          e.preventDefault()
        });
        document.getElementsByTagName("canvas")[0].addEventListener('touchmove', (e) => {
          // console.log(e.type)
          if (e.targetTouches[1]) {
            if (this.distance !== null) {
              let x1=e.targetTouches[0].clientX
              let x2=e.targetTouches[1].clientX
              let y1=e.targetTouches[0].clientY
              let y2=e.targetTouches[1].clientY
              let dx = Math.abs(x1-x2) 
              let dy = Math.abs(y1-y2)
              let dis =Math.sqrt((dx * dx) + (dy * dy))
              // console.log(dis,"dis")
              // console.log(dis-this.distance ,"chazhi")
              this.graph.zoom((dis-this.distance)>0?1.01:0.99,{ x: Math.min(x1,x2)+(dx/2), y: Math.min(y1,y2)+(dy/2) })
              this.distance=dis
            }else{
              let dx = Math.abs(e.targetTouches[0].clientX - e.targetTouches[1].clientX) 
              let dy = Math.abs(e.targetTouches[0].clientY - e.targetTouches[1].clientY)
              this.distance =Math.sqrt((dx * dx) + (dy * dy))
            }
          } else {
            this.distance = null
          }
          if (this.clientY !== null && this.clientX !== null) {
            // if (this.graph) {
            //   let dx = e.targetTouches[0].clientX - this.clientX
            //   let dy = e.targetTouches[0].clientY - this.clientY
            //   // console.log(dx, dy)
            //   this.graph.translate(dx, dy)
            // }
            var evt = new MouseEvent("mousemove", {
              bubbles: true,
              cancelable: true,
              view: window,
              screenX: e.targetTouches[0].clientX,
              screenY: e.targetTouches[0].clientY,
              clientX: e.targetTouches[0].clientX,
              clientY: e.targetTouches[0].clientY,
              ctrlKey: false,
              altKey: false,
              shiftKey: false,
              metaKey: false,
              button: 0,
              buttons: 1,
            });
            document.getElementsByTagName("canvas")[0].dispatchEvent(evt)
            // console.log(evt)
            this.clientX = e.targetTouches[0].clientX
            this.clientY = e.targetTouches[0].clientY
          }
          e.preventDefault()
        });

      },
      recolor() {
        if (this.isStart) {
          return
        }
        this.data = this.graph.getNodes();
        this.data.forEach(n => {
          let node = {
            id: n.defaultCfg.model.id,
            label: n.defaultCfg.model.model,
            style: {
              fill: '#409EFF',
            },
          };
          this.graph.updateItem(n, node)
        });
      },
      tryShowAdd() {
        if (this.isStart) {
          return
        }
        this.showAdd = true
      },
      delFn(evt) {
        let item = evt.item; // 被操作的节点 item
        this.graph.removeItem(item);
      },
      connFn(evt) {
        let item = evt.item; // 被操作的节点 item
        if (this.tempNode === null) {
          this.tempNode = item
          this.hint = "请点击要连接的第2个节点,如需取消请点击[Connect]"
          return
        }
        if (this.tempNode.defaultCfg.id !== item.defaultCfg.id) {
          let edge = {
            source: this.tempNode,
            target: item,
          }
          this.graph.addItem('edge', edge)
          this.graph.refresh();
          this.tempNode = null
          this.hint = "连接成功!如需继续请点击要连接的第1个节点,如需取消请点击[Connect]"
          return
        }
      },
      dfs(n) {
        // console.log(this.data[n].defaultCfg.id)
        // console.log(this.data[n])
        this.step++
        setTimeout(() => {
          this.setForward(this.data[n])
        }, 500 * this.step);
        // console.log(this.data[n].next)
        for (let i = 0; i < this.data[n].next.length; i++) {
          let t = this.getIndex(this.data[n].next[i])
          if (!this.data[t].read) {
            this.data[t].read = true
            this.dfs(t)
            // this.data[t].read = false
            this.step++
            setTimeout(() => {
              this.setBackward(this.data[t])
            }, 500 * this.step);
          }
        }
      },
      getIndex(id) {
        let index = 0
        for (let i = 0; i < this.data.length; i++) {
          if (this.data[i].defaultCfg.id === id) {
            index = i
          }
        }
        return index
      },
      realStart(evt) {
        if (this.isStart) {
          return
        }
        console.log(this.is)
        this.isStart = true
        this.graph.off('node:click', this.realStart)
        this.hint = "运行中"
        let item = evt.item; // 被操作的节点 item
        this.data = this.graph.getNodes();
        this.data.forEach(n => {
          n.next = []
          n.read = false
          n.getEdges().forEach(e => {
            if (e.defaultCfg.sourceNode.defaultCfg.id !== n.defaultCfg.id) {
              n.next.push(e.defaultCfg.sourceNode.defaultCfg.id)
            } else {
              n.next.push(e.defaultCfg.targetNode.defaultCfg.id)
            }
          })
          n.next.sort(function (a, b) {
            return Number(a) - Number(b)
          })
        });
        this.data.sort(function (a, b) {
          return Number(a.defaultCfg.id) - Number(b.defaultCfg.id)
        })
        let begin = this.getIndex(item.defaultCfg.id)
        // console.log(begin, "begin")
        // console.log(this.data, "data")
        this.step = 0
        this.data[begin].read = true
        this.dfs(begin)
        this.step++
        setTimeout(() => {
          this.setBackward(this.data[begin])
          this.isStart = false
          this.isReady = false
          this.isConnect = false
          this.isDelete = false
          this.hint = ""
        }, 500 * this.step);
        this.step = 0
      },
      switchConnect() {
        if (this.isStart) {
          return
        }
        if (this.isReady) {
          this.switchReady()
        }
        if (this.isDelete) {
          this.switchDelete()
        }
        this.isConnect = !this.isConnect
        if (this.isConnect) {
          this.graph.on('node:click', this.connFn);
          this.hint = "请点击要连接的第1个节点,如需取消请点击[Connect]"
        } else {
          this.graph.off('node:click', this.connFn);
          this.tempNode = null
          this.hint = ""
        }
      },
      switchDelete() {
        if (this.isStart) {
          return
        }
        if (this.isReady) {
          this.switchReady()
        }
        if (this.isConnect) {
          this.switchConnect()
        }
        this.isDelete = !this.isDelete
        if (this.isDelete) {
          this.graph.on('node:click', this.delFn);
          this.graph.on('edge:click', this.delFn);
          this.hint = "请点击要删除的节点或边,如需取消请点击[Delete]"
        } else {
          this.graph.off('node:click', this.delFn);
          this.graph.off('edge:click', this.delFn);
          this.hint = ""
        }
      },
      deleteAll() {
        if (this.isStart) {
          return
        }
        this.graph.clear();
        this.hint = "清除成功"
      },
      addNode() {
        if (this.numToAdd === "") {
          this.hint = "添加的数字不能为空"
          this.showAdd = false
          return
        }
        try {
          if (this.numToAdd[0] === "[" && this.numToAdd.charAt(this.numToAdd.length - 1) === "]" && this.numToAdd
            .split(
              "[").length == 2) {
            let arr = eval("(" + this.numToAdd + ")")
            for (let i = 0; i < arr.length; i++) {
              if (!isNaN(arr[i])) {
                let node = {
                  id: arr[i] + "",
                  label: arr[i] + "",
                  address: 'cq',
                  x: Math.floor((Math.random() * 300) + 100),
                  y: Math.floor((Math.random() * 300) + 100),
                };
                this.graph.addItem('node', node)
              }
            }
            this.graph.refresh();
            this.showAdd = false
            return
          }
          let num = eval("(" + this.numToAdd + ")")
          if (!isNaN(num)) {
            let node = {
              id: num + "",
              label: num + "",
              address: 'cq',
              x: Math.floor((Math.random() * 300) + 100),
              y: Math.floor((Math.random() * 300) + 100),
            };
            this.graph.addItem('node', node)
            this.graph.refresh();
            this.showAdd = false
            return
          }
          this.hint = "输入数据无效"
          this.showAdd = false
        } catch (err) {
          this.hint = "输入数据无效"
          this.showAdd = false
        } finally {
          this.numToAdd = ""
        }


      },
      initG6() {
        this.data = {
          // 点集
          nodes: [{
              id: '0',
              label: '0', // 边的文本
              next: [1],
              read: false
            },
            {
              id: '1',
              label: '1', // 边的文本
              next: [2, 3],
              read: false
            },
            {
              id: '2',
              label: '2', // 边的文本
              next: [4],
              read: false
            },
            {
              id: '3',
              label: '3', // 边的文本
              next: [4, 5, 7],
              read: false
            },
            {
              id: '4',
              label: '4', // 边的文本
              next: [2, 3, 8],
              read: false
            },
            {
              id: '5',
              label: '5', // 边的文本
              next: [4, 3, 6],
              read: false
            },
            {
              id: '6',
              label: '6', // 边的文本
              next: [5],
              read: false
            },
            {
              id: '7',
              label: '7', // 边的文本
              next: [3, 9, 10],
              read: false
            },
            {
              id: '8',
              label: '8', // 边的文本
              next: [4],
              read: false
            },
            {
              id: '9',
              label: '9', // 边的文本
              next: [7],
              read: false
            },
            {
              id: '10',
              label: '10', // 边的文本
              next: [7],
              read: false
            },
          ],
          // 边集
          edges: [
            // 表示一条从 node1 节点连接到 node2 节点的边
            {
              source: '1',
              target: '2',
            },
            {
              source: '1',
              target: '3',
            },
            {
              source: '2',
              target: '4',
            },
            {
              source: '3',
              target: '4',
            },
            {
              source: '4',
              target: '5',
            },
            {
              source: '4',
              target: '8',
            },
            {
              source: '5',
              target: '3',
            },
            {
              source: '5',
              target: '6',
            },
            {
              source: '3',
              target: '7',
            },
            {
              source: '7',
              target: '9',
            },
            {
              source: '7',
              target: '10',
            },
            {
              source: '0',
              target: '1',
            },
          ],
        };
        this.grid = new G6.Grid();
        this.graph = new G6.Graph({
          container: 'g6', // 指定图画布的容器 id，与第 9 行的容器对应
          // 画布宽高
          width: window.screen.width,
          height: window.innerHeight,
          animate: true,
          // fitView:true,
          layout: {
            type: 'radial',
            preventOverlap: true,
            linkDistance: 250, // 指定边距离为100
          },
          plugins: [this.grid],
          modes: {
            default: ['drag-canvas', 'zoom-canvas', 'drag-node'], // 允许拖拽画布、放缩画布、拖拽节点
          },
          defaultEdge: {
            size: 4, // 节点大小
            style: {
              fill: '#409EFF', // 节点填充色
              opacity: 1,
              stroke: '#E4E7ED', // 边描边颜色
            },
          },
          defaultNode: {
            size: 45, // 节点大小
            // ...                 // 节点的其他配置
            // 节点样式配置
            style: {
              fill: '#409EFF', // 节点填充色
              stroke: '#DCDFE6', // 节点描边色
              lineWidth: 1, // 节点描边粗细
            },
            // 节点上的标签文本配置
            labelCfg: {
              // 节点上的标签文本样式配置
              style: {
                fill: '#fff', // 节点标签文字颜色
              },
            },
          },
          nodeStateStyles: {
            // 鼠标 hover 上节点，即 hover 状态为 true 时的样式
            hover: {
              fill: 'lightsteelblue',
            },

            // 鼠标点击节点，即 click 状态为 true 时的样式
            click: {
              stroke: '#000',
              lineWidth: 3,
            },

          },
          // 边不同状态下的样式集合
          edgeStateStyles: {
            // 鼠标点击边，即 click 状态为 true 时的样式
            click: {
              stroke: 'steelblue',
            },
          },
        });
        this.graph.fitView(20)
        this.graph.read(this.data);
      },
      switchReady() {
        if (this.isDelete) {
          this.switchDelete()
        }
        if (this.isConnect) {
          this.switchConnect()
        }
        if (this.isReady) {
          this.graph.off('node:click', this.realStart)
          this.hint = ""
          this.isReady = false
        } else {
          this.graph.on('node:click', this.realStart)
          this.hint = "请点击需要起始的节点，如需取消请点击[Start]"
          this.isReady = true
        }
      },
      setBackward(item) {
        let node = {
          id: item.defaultCfg.model.id,
          label: item.defaultCfg.model.label,
          style: {
            fill: '#DCDFE6',
          },
        };
        this.graph.updateItem(item, node)
      },
      setForward(item) {
        let node = {
          id: item.defaultCfg.model.id,
          label: item.defaultCfg.model.label,
          style: {
            fill: '#F56C6C',
          },
        };
        this.graph.updateItem(item, node)
      },

    },
    components: {},
  }
</script>

<style>
  .button-group {
    position: absolute;
    bottom: 20px;
    width: 100%;
    z-index: 100;
  }

  .flex-contain {
    padding: 0 auto
  }

  .start {
    background-color: #C0C4CC !important;
    opacity: 0;
    cursor: default !important;
  }

  .big-circle {
    height: 100px;
    width: 100px;
    line-height: 100px;
    border-radius: 50px;
    color: green;
  }

  .small-circle {
    height: 50px;
    width: 50px;
    line-height: 50px;
    border-radius: 25px;
    color: green;
    font-size: 10px;
  }

  #g6 {
    /* background:red; */
    /* z-index:1000; */
    position: relative;
  }

  .flex-demo {
    margin: 0 auto;
    text-align: center;
    color: green;
    background-color: #fff;
    background-clip: padding-box;
    cursor: pointer
  }
</style>
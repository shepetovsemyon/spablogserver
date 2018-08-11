// <reference path ="./node_modules/@types/jquery/index.d.ts"/> 

//import * as $ from 'jquery'; 

class Loader{
    private filename = "";

    private template: string;
    private cbList: [((string) => void), boolean][] = [];
    private isLoading = false;

    private clearCBList(){
        for (let i = this.cbList.length - 1; i > 0; i--) {
            if(!this.cbList[i][1]) this.cbList.splice(i,1);          
        }
    }

    public Load(callback: (string: string) => void) {
        if(this.template == null && !this.isLoading){
            this.isLoading = true;

            this.cbList.push([callback, true]);
            
            $.get(this.filename, (data) => {
                this.isLoading = false;
                this.template = data;
                
                this.cbList.forEach(element => {
                    if(element[1]){
                        element[0](data);
                        element[1] = false;
                    }
                });

                this.clearCBList();
            });
            return;
        }
        
        if(this.isLoading){
            this.cbList.push([callback, true]);
        }else{
            callback(this.template);
        }        
    }

    constructor(filename: string) {
        this.filename = filename;
    }   
}

interface Article{
    Title: string;
    Content: string;
}

class ArticleManager {
    private tLoader = new Loader("templates/article.html");

    private _node = <HTMLSpanElement>(document.createElement("span"));

    private LoadData(callback : (Article: Article) => void){
        $.get('/api',
            (data) => {                
                let article: Article = <Article>JSON.parse(data);
                callback(article);            
            });
    }

    private _render(node: HTMLElement, template, title, content : string){
        this._node.innerHTML = template; 
            if(this._node.childNodes.length < 1)return;

            this._node.getElementsByClassName("templ-article-header")[0]
                .innerHTML = title;  

            this._node.getElementsByClassName("templ-article-context")[0]
            .innerHTML = content;

            node.appendChild(this._node.childNodes[0]);
    }

    public Render(node: HTMLElement, title = "", content = ""): void {
        this.tLoader.Load((templ) => {
            this._render(node, templ, title, content);
        });        
    } 

    public Render1(node: HTMLElement): void {
        this.tLoader.Load((templ) => {
            this.LoadData((article) => {                  
                console.log(article);              
                this._render(node, templ, article.Title, article.Content);
            });
        });        
    } 
}

class BlogRendered{
    private dataManager: ArticleManager;

    constructor(dataManager: ArticleManager){
        this.dataManager = dataManager;
    }

    public async RenderAll(){
        let blogs = document.body.getElementsByTagName("Blog");

        for(let i = 0; i < blogs.length; i++){
            await this.Render(<HTMLElement>blogs[i]);
        }
    }
    
    
    async Render(blog: HTMLElement ){
        
        let descrResp = await fetch('/api/blog');

        let descrJson = await descrResp.json();

        let descr : BlogDescr = <BlogDescr>descrJson;//JSON.parse(descrJson);
        
        console.log(descr);

        let count = 10;
        let url = `/api/blog/articles?offset=${descr.Count - count}&count=${count}`

        let articlesResp = await fetch(url)

        console.log(articlesResp);

        let articlesJson = await articlesResp.json();

        console.log(articlesResp);

        let articles : Article[] = <Article[]>articlesJson;//JSON.parse(articlesJson);
        
        //return new Promiss(articles);
        //this.dataManager.Render

        this.RenderSync(blog, articles);
    }

    private RenderSync(blog: HTMLElement, articles: Article[] ){

        articles.forEach(article => {            
            let div = document.createElement('Div');
            blog.appendChild(div);
            this.dataManager.Render(div, article.Title, article.Content)
        });

    }    
}

interface BlogDescr {
    Autor: string;
    Count: number
}

class Blog{
    private Descr: BlogDescr;

    //ma

}

$(document).ready(() => {
    let a = new ArticleManager();

    let b = new BlogRendered(a);

    b.RenderAll();

    a.Render($("body").get(0), "ABC1", "abcabc");
    a.Render($("body").get(0), "TEXT1", "text text");
   // a.Render1($("body").get(0));
});


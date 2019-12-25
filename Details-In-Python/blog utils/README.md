## formatter.py
  Markdown like this
  ```html
  ...
  <div>
	<style>
		coding
	</style>
	<other label>
		coding 
	</other label>
  </div>
  ...
  <div>
	<style>
		coding
	</style>
		coding
	<other label>
		coding 
	</other label>
  </div>
  ...
  ```
  Since html code in markdown is parsed by hexo can generate lots of ``</br>``, many blank lines, use this script to get like this 
  ```html
  ...
  <div>		<other label>		coding 	</other label>  </div>
  ...
  <div>			coding	<other label>		coding 	</other label>  </div>
  ...
  ```
  <style> label is removed since it is useless.
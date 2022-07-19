#409. 最长回文串

##题目描述
<p>给定一个包含大写字母和小写字母的字符串<meta charset="UTF-8" />&nbsp;<code>s</code>&nbsp;，返回&nbsp;<em>通过这些字母构造成的 <strong>最长的回文串</strong></em>&nbsp;。</p>

<p>在构造过程中，请注意 <strong>区分大小写</strong> 。比如&nbsp;<code>"Aa"</code>&nbsp;不能当做一个回文字符串。</p>

<p>&nbsp;</p>

<p><strong>示例 1: </strong></p>

<pre>
<strong>输入:</strong>s = "abccccdd"
<strong>输出:</strong>7
<strong>解释:</strong>
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong>s = "a"
<strong>输入:</strong>1
</pre>

<p><strong>示例 3:</strong></p>

<pre>
<strong>输入:</strong>s = "bb"
<strong>输入:</strong> 2
</pre>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ul>
	<li><code>1 &lt;= s.length &lt;= 2000</code></li>
	<li><code>s</code>&nbsp;只能由小写和/或大写英文字母组成</li>
</ul>

##题解
“回文串”是一个正读和反读都一样的字符串
1. 使用一个`map`记录所有字符出现的数量，一个字符出现偶数次即可回文数中的一对，一个字符出现奇数次时，有一次必须放在中间位置
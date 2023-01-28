
# uniar

`uniar` is database and management your scene collections CLI tool for [UNI'S ON AIR](https://keyahina-unisonair.com/).

## Usage		

```
uniar is UNI'S ON AIR music and scene cards database and manage your scene cards collection tool.

Usage:
  uniar [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  doc         Generate uniar document
  help        Help about any command
  list        Show data
  regist      Regist data
  server      Server
  setup       Setup uniar

Flags:
      --config string   config file (default is $HOME/.uniar.yaml)
  -h, --help            help for uniar
  -v, --version         version for uniar

Use "uniar [command] --help" for more information about a command.

```

## Install

```
$ brew tap litencatt/tap
$ brew install litencatt/tap/uniar
```

## Commands

### List

```
$ uniar list
uniar is UNI'S ON AIR music and scene cards database and manage your scene cards collection tool.

Usage:
  uniar [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  doc         Generate uniar document
  help        Help about any command
  list        Show data
  regist      Regist data
  server      Server
  setup       Setup uniar

Flags:
      --config string   config file (default is $HOME/.uniar.yaml)
  -h, --help            help for uniar
  -v, --version         version for uniar

Use "uniar [command] --help" for more information about a command.

```

### Usage

```
$ uniar list scene -h
Show scene card list

Usage:
  uniar list scene [flags]

Aliases:
  scene, s

Flags:
  -c, --color string            Color filter(e.g. -c Red or -c r)
  -d, --detail                  Show detail
  -f, --full-name               Show pohtograph full name
      --have                    Show only scenes you have
  -h, --help                    help for scene
  -i, --ignore-columns string   Ignore columns to display(VoDa50,DaPe50,...)
  -m, --member string           Member filter(e.g. -m 加藤史帆)
  -n, --not-have                Show only scenes you NOT have
  -p, --photograph string       Photograph filter(e.g. -p JOYFULLOVE)
  -s, --sort string             Sort target rank.(all35, voda50, ...)

Global Flags:
      --config string   config file (default is $HOME/.uniar.yaml)

```

```
$ uniar list scene -f | head
+--------+------------------------------------+---------+------------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| COLOR  |             PHOTOGRAPH             | SSRPLUS |   MEMBER   | EXPECT | TOTAL | ALL35 | VODA50 | DAPE50 | VOPE50 | VO85 | DA85 | PE85 | HAVE  |
+--------+------------------------------------+---------+------------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| Blue   | 青春の馬                           | true    | 濱岸ひより |   3.72 | 13077 |     1 |     31 |      2 |      2 |   35 |   44 |    2 | false |
| Purple | なぜ　恋をして来なかったんだろう？ | true    | 藤吉夏鈴   |   3.72 | 13076 |     2 |     32 |      1 |      1 |   36 |   45 |    1 | false |
| Yellow | 誰がその鐘を鳴らすのか？           | true    | 山﨑天     |   3.68 | 13076 |     3 |      3 |     22 |      4 |    1 |   19 |   15 | false |
| Green  | こんなに好きになっちゃっていいの？ | true    | 加藤史帆   |   3.68 | 13067 |     4 |      2 |     23 |      5 |    2 |   18 |   17 | false |
| Blue   | Nobody’s fault                    | true    | 渡邉理佐   |   3.68 | 13065 |     5 |      1 |      3 |     37 |   39 |    1 |   45 | false |
| Red    | 誰がその鐘を鳴らすのか？           | true    | 菅井友香   |   3.68 | 13064 |     6 |     20 |      6 |      6 |   11 |   17 |    7 | false |
| Purple | こんなに好きになっちゃっていいの？ | true    | 東村芽依   |   3.68 | 13059 |     7 |      5 |      5 |     19 |   12 |    2 |   16 | false |

```

```
$ uniar list scene -f -c Blue
+-------+------------------------------------+---------+------------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| COLOR |             PHOTOGRAPH             | SSRPLUS |   MEMBER   | EXPECT | TOTAL | ALL35 | VODA50 | DAPE50 | VOPE50 | VO85 | DA85 | PE85 | HAVE  |
+-------+------------------------------------+---------+------------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| Blue  | 青春の馬                           | true    | 濱岸ひより |   3.72 | 13077 |     1 |      6 |      1 |      1 |    4 |    7 |    1 | false |
| Blue  | Nobody’s fault                    | true    | 渡邉理佐   |   3.68 | 13065 |     2 |      1 |      2 |      6 |    5 |    1 |    8 | false |
| Blue  | アンビバレント                     | true    | 田村保乃   |   3.68 | 12976 |     3 |      5 |      3 |      4 |    6 |    3 |    2 | false |
| Blue  | ドレミソラシド                     | true    | 森本茉莉   |   3.68 | 12974 |     4 |      2 |      5 |      2 |    1 |    4 |    3 | false |
| Blue  | 風に吹かれても                     | true    | 上村莉菜   |   3.68 | 12931 |     5 |      3 |      6 |      3 |    2 |    5 |    4 | false |
| Blue  | ガラスを割れ！                     | true    | 土生瑞穂   |   3.68 | 12819 |     6 |      4 |      4 |      5 |    3 |    2 |    7 | false |
| Blue  | UNI'S ON AIR 3rd ANNIVERSARY       | false   | 潮紗理菜   |   3.68 | 11435 |     7 |      7 |     14 |     11 |    7 |   10 |   29 | false |

```

```
$ uniar list scene -f -m 加藤史帆
+--------+------------------------------------+---------+----------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| COLOR  |             PHOTOGRAPH             | SSRPLUS |  MEMBER  | EXPECT | TOTAL | ALL35 | VODA50 | DAPE50 | VOPE50 | VO85 | DA85 | PE85 | HAVE  |
+--------+------------------------------------+---------+----------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| Green  | こんなに好きになっちゃっていいの？ | true    | 加藤史帆 |   3.68 | 13067 |     1 |      1 |      1 |      1 |    1 |    1 |    1 | false |
| Yellow | 君しか勝たん                       | false   | 加藤史帆 |   3.68 | 11083 |     2 |      2 |      2 |      4 |    4 |    2 |    5 | false |
| Green  | こんなに好きになっちゃっていいの？ | false   | 加藤史帆 |   3.68 | 10817 |     3 |      3 |      5 |      2 |    2 |    5 |    4 | false |
| Blue   | NEW YEAR'23                        | false   | 加藤史帆 |   3.49 | 11269 |     4 |      4 |      4 |      3 |    3 |    4 |    3 | false |
| Purple | 清夏彩る涼の装                     | false   | 加藤史帆 |    3.4 | 11357 |     5 |      5 |      3 |      5 |    5 |    3 |    2 | false |
| Red    | ハッピーオーラ                     | true    | 加藤史帆 |   2.46 | 12932 |     6 |      6 |      6 |      6 |    6 |    7 |    6 | false |
| Blue   | キュン                             | true    | 加藤史帆 |   2.46 | 12560 |     7 |      7 |      7 |      7 |    7 |    8 |    9 | false |

```

```
$ uniar list scene -f -p キュン
+--------+------------+---------+------------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| COLOR  | PHOTOGRAPH | SSRPLUS |   MEMBER   | EXPECT | TOTAL | ALL35 | VODA50 | DAPE50 | VOPE50 | VO85 | DA85 | PE85 | HAVE  |
+--------+------------+---------+------------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| Green  | キュン     | true    | 小坂菜緒   |   3.68 | 12582 |     1 |      1 |      1 |      1 |    1 |    1 |    1 | false |
| Green  | キュン     | false   | 小坂菜緒   |   3.68 | 10332 |     2 |      2 |      2 |      2 |    2 |    2 |    2 | false |
| Red    | キュン     | true    | 松田好花   |   2.75 | 12902 |     3 |      3 |      3 |      3 |    3 |    3 |    3 | false |
| Purple | キュン     | true    | 上村ひなの |   2.52 | 12768 |     4 |      4 |      4 |      5 |    5 |    4 |    4 | false |
| Green  | キュン     | true    | 佐々木美玲 |   2.52 | 12599 |     5 |      5 |      5 |      4 |    4 |    6 |    5 | false |
| Blue   | キュン     | true    | 加藤史帆   |   2.46 | 12560 |     6 |      6 |      6 |      6 |    6 |    5 |    6 | false |
| Purple | キュン     | true    | 潮紗理菜   |   2.32 | 12647 |     7 |      7 |      7 |      7 |    7 |    7 |    7 | false |

```

```
$ uniar list scene -f -c Blue -m 加藤史帆 -p キュン
+-------+------------+---------+----------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| COLOR | PHOTOGRAPH | SSRPLUS |  MEMBER  | EXPECT | TOTAL | ALL35 | VODA50 | DAPE50 | VOPE50 | VO85 | DA85 | PE85 | HAVE  |
+-------+------------+---------+----------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| Blue  | キュン     | true    | 加藤史帆 |   2.46 | 12560 |     1 |      1 |      1 |      1 |    1 |    1 |    1 | false |
| Blue  | キュン     | false   | 加藤史帆 |   2.46 | 10310 |     2 |      2 |      2 |      2 |    2 |    2 |    2 | false |
+-------+------------+---------+----------+--------+-------+-------+--------+--------+--------+------+------+------+-------+

```

```
$ uniar list scene --ignore-columns All35,VoDa50,DaPe50,VoPe50 | head
+--------+--------------------------------+---------+------------+--------+-------+------+------+------+-------+
| COLOR  |           PHOTOGRAPH           | SSRPLUS |   MEMBER   | EXPECT | TOTAL | VO85 | DA85 | PE85 | HAVE  |
+--------+--------------------------------+---------+------------+--------+-------+------+------+------+-------+
| Blue   | 馬                             | true    | 濱岸ひより |   3.72 | 13077 |   35 |   44 |    2 | false |
| Purple | なぜ恋                         | true    | 藤吉夏鈴   |   3.72 | 13076 |   36 |   45 |    1 | false |
| Yellow | 誰鐘                           | true    | 山﨑天     |   3.68 | 13076 |    1 |   19 |   15 | false |
| Green  | こん好き                       | true    | 加藤史帆   |   3.68 | 13067 |    2 |   18 |   17 | false |
| Blue   | ノバフォ                       | true    | 渡邉理佐   |   3.68 | 13065 |   39 |    1 |   45 | false |
| Red    | 誰鐘                           | true    | 菅井友香   |   3.68 | 13064 |   11 |   17 |    7 | false |
| Purple | こん好き                       | true    | 東村芽依   |   3.68 | 13059 |   12 |    2 |   16 | false |

```

```
$ uniar list scene -d --ignore-columns All35,VoDa50,DaPe50,VoPe50 | head
+--------+--------------------------------+---------+------------+--------+-------+------+------+------+------+------+------+-------+
| COLOR  |           PHOTOGRAPH           | SSRPLUS |   MEMBER   | EXPECT | TOTAL | VO85 | DA85 | PE85 |  VO  |  DA  |  PE  | HAVE  |
+--------+--------------------------------+---------+------------+--------+-------+------+------+------+------+------+------+-------+
| Blue   | 馬                             | true    | 濱岸ひより |   3.72 | 13077 |   35 |   44 |    2 | 2989 | 3015 | 6643 | false |
| Purple | なぜ恋                         | true    | 藤吉夏鈴   |   3.72 | 13076 |   36 |   45 |    1 | 2971 | 2974 | 6701 | false |
| Yellow | 誰鐘                           | true    | 山﨑天     |   3.68 | 13076 |    1 |   19 |   15 | 5092 | 3756 | 3798 | false |
| Green  | こん好き                       | true    | 加藤史帆   |   3.68 | 13067 |    2 |   18 |   17 | 5102 | 3780 | 3755 | false |
| Blue   | ノバフォ                       | true    | 渡邉理佐   |   3.68 | 13065 |   39 |    1 |   45 | 3020 | 6614 | 3001 | false |
| Red    | 誰鐘                           | true    | 菅井友香   |   3.68 | 13064 |   11 |   17 |    7 | 3779 | 3888 | 4967 | false |
| Purple | こん好き                       | true    | 東村芽依   |   3.68 | 13059 |   12 |    2 |   16 | 3745 | 5066 | 3818 | false |

```

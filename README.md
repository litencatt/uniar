
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
| Green  | こんなに好きになっちゃっていいの？ | true    | 加藤史帆   |   3.68 | 13067 |     1 |      1 |      2 |      1 |    1 |    2 |    2 | false |
| Purple | Buddies                            | true    | 山﨑天     |   3.72 | 13323 |     2 |      2 |      1 |      6 |    4 |    1 |    6 | false |
| Yellow | 誰がその鐘を鳴らすのか？           | true    | 山﨑天     |   3.68 | 13076 |     3 |      3 |      4 |      3 |    2 |    6 |    5 | false |
| Red    | ハッピーオーラ                     | true    | 松田好花   |   3.68 | 12974 |     4 |      4 |      7 |      2 |    3 |   13 |    4 | false |
| Yellow | ハッピーオーラ                     | true    | 高瀬愛奈   |   3.68 | 12972 |     5 |      7 |      5 |      5 |    8 |   10 |    3 | false |
| Red    | 二人セゾン                         | true    | 小池美波   |   3.68 | 12859 |     6 |      5 |      6 |      7 |    7 |    7 |   11 | false |
| Red    | アザトカワイイ                     | true    | 佐々木美玲 |   3.72 | 13324 |     7 |     18 |      3 |      4 |   19 |   20 |    1 | false |

```

```
$ uniar list scene -f -c Blue
+-------+------------------------------------+---------+------------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| COLOR |             PHOTOGRAPH             | SSRPLUS |   MEMBER   | EXPECT | TOTAL | ALL35 | VODA50 | DAPE50 | VOPE50 | VO85 | DA85 | PE85 | HAVE  |
+-------+------------------------------------+---------+------------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| Blue  | 風に吹かれても                     | true    | 上村莉菜   |   3.68 | 12931 |     1 |      1 |      2 |      1 |    1 |    2 |    2 | false |
| Blue  | Nobody’s fault                    | true    | 渡邉理佐   |   3.68 | 13065 |     2 |      2 |      1 |      8 |    9 |    1 |   11 | false |
| Blue  | アンビバレント                     | true    | 田村保乃   |   3.68 | 12976 |     3 |      7 |      3 |      3 |   12 |    6 |    1 | false |
| Blue  | ガラスを割れ！                     | true    | 土生瑞穂   |   3.68 | 12819 |     4 |      4 |      4 |      5 |    8 |    4 |   10 | false |
| Blue  | NEW YEAR'23                        | false   | 加藤史帆   |   3.49 | 11269 |     5 |      6 |      5 |      4 |    6 |    7 |    8 | false |
| Blue  | UNI'S ON AIR 1st ANNIVERSARY       | false   | 山﨑天     |   3.68 | 11119 |     6 |      9 |      8 |      2 |    2 |   15 |    5 | false |
| Blue  | UNI'S ON AIR 3rd ANNIVERSARY       | false   | 潮紗理菜   |   3.68 | 11435 |     7 |      3 |     10 |      9 |    3 |    5 |   17 | false |

```

```
$ uniar list scene -f -m 加藤史帆
+--------+------------------------------------+---------+----------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| COLOR  |             PHOTOGRAPH             | SSRPLUS |  MEMBER  | EXPECT | TOTAL | ALL35 | VODA50 | DAPE50 | VOPE50 | VO85 | DA85 | PE85 | HAVE  |
+--------+------------------------------------+---------+----------+--------+-------+-------+--------+--------+--------+------+------+------+-------+
| Green  | こんなに好きになっちゃっていいの？ | true    | 加藤史帆 |   3.68 | 13067 |     1 |      1 |      1 |      1 |    1 |    1 |    1 | false |
| Yellow | 君しか勝たん                       | false   | 加藤史帆 |   3.68 | 11083 |     2 |      2 |      2 |      4 |    4 |    2 |    5 | false |
| Green  | こんなに好きになっちゃっていいの？ | false   | 加藤史帆 |   3.68 | 10817 |     3 |      3 |      4 |      2 |    2 |    4 |    2 | false |
| Blue   | NEW YEAR'23                        | false   | 加藤史帆 |   3.49 | 11269 |     4 |      4 |      5 |      3 |    3 |    5 |    4 | false |
| Purple | 清夏彩る涼の装                     | false   | 加藤史帆 |    3.4 | 11357 |     5 |      5 |      3 |      5 |    5 |    3 |    3 | false |
| Red    | ハッピーオーラ                     | true    | 加藤史帆 |   2.46 | 12932 |     6 |      6 |      6 |      6 |    6 |    7 |    7 | false |
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
| Blue   | キュン     | true    | 加藤史帆   |   2.46 | 12560 |     4 |      4 |      4 |      4 |    4 |    4 |    4 | false |
| Purple | キュン     | true    | 上村ひなの |   2.52 | 12768 |     5 |      5 |      5 |      5 |    5 |    5 |    5 | false |
| Green  | キュン     | true    | 佐々木美玲 |   2.52 | 12599 |     6 |      6 |      6 |      6 |    6 |    6 |    6 | false |
| Purple | キュン     | true    | 潮紗理菜   |   2.32 | 12647 |     7 |      7 |      7 |      7 |    7 |    8 |    7 | false |

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
| Green  | こん好き                       | true    | 加藤史帆   |   3.68 | 13067 |    1 |    2 |    2 | false |
| Purple | Buddies                        | true    | 山﨑天     |   3.72 | 13323 |    4 |    1 |    6 | false |
| Yellow | 誰鐘                           | true    | 山﨑天     |   3.68 | 13076 |    2 |    6 |    5 | false |
| Red    | ハピオラ                       | true    | 松田好花   |   3.68 | 12974 |    3 |   13 |    4 | false |
| Yellow | ハピオラ                       | true    | 高瀬愛奈   |   3.68 | 12972 |    8 |   10 |    3 | false |
| Red    | セゾン                         | true    | 小池美波   |   3.68 | 12859 |    7 |    7 |   11 | false |
| Red    | アザカワ                       | true    | 佐々木美玲 |   3.72 | 13324 |   19 |   20 |    1 | false |

```

```
$ uniar list scene -d --ignore-columns All35,VoDa50,DaPe50,VoPe50 | head
+--------+--------------------------------+---------+------------+--------+-------+------+------+------+------+------+------+-------+
| COLOR  |           PHOTOGRAPH           | SSRPLUS |   MEMBER   | EXPECT | TOTAL | VO85 | DA85 | PE85 |  VO  |  DA  |  PE  | HAVE  |
+--------+--------------------------------+---------+------------+--------+-------+------+------+------+------+------+------+-------+
| Green  | こん好き                       | true    | 加藤史帆   |   3.68 | 13067 |    1 |    2 |    2 | 5102 | 3780 | 3755 | false |
| Purple | Buddies                        | true    | 山﨑天     |   3.72 | 13323 |    4 |    1 |    6 | 3033 | 6795 | 3065 | false |
| Yellow | 誰鐘                           | true    | 山﨑天     |   3.68 | 13076 |    2 |    6 |    5 | 5092 | 3756 | 3798 | false |
| Red    | ハピオラ                       | true    | 松田好花   |   3.68 | 12974 |    3 |   13 |    4 | 5059 | 2920 | 4565 | false |
| Yellow | ハピオラ                       | true    | 高瀬愛奈   |   3.68 | 12972 |    8 |   10 |    3 | 3761 | 3820 | 4961 | false |
| Red    | セゾン                         | true    | 小池美波   |   3.68 | 12859 |    7 |    7 |   11 | 4006 | 4578 | 3845 | false |
| Red    | アザカワ                       | true    | 佐々木美玲 |   3.72 | 13324 |   19 |   20 |    1 | 3023 | 3036 | 6835 | false |

```
